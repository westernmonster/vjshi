package parser

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func GrabRecentSales() (err error) {
	url := "https://www.vjshi.com/ranking/sales"
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // for debug use
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("enable-automation", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 20*time.Second)
	defer cancel()

	var htmlContent string
	err = chromedp.Run(timeoutCtx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, errAc := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(ctx)
			if errAc != nil {
				return errAc
			}
			return nil
		}),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`.nc_1_nocaptcha`, chromedp.ByQuery),
		dragElement("#nc_1__bg"),
		chromedp.WaitVisible(`.video-card__list`, chromedp.ByQuery),
		chromedp.EvaluateAsDevTools(`JSON.stringify({name: "a"})`, &htmlContent),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		fmt.Printf("error log: %+v", err)
		return
	}

	fmt.Println(htmlContent)

	return
}

func dragElement(sel interface{}) chromedp.QueryAction {
	return chromedp.QueryAfter(sel, func(ctx context.Context, id runtime.ExecutionContextID, node ...*cdp.Node) error {
		if len(node) == 0 {
			return fmt.Errorf("Can not find this node in content")
		}
		return mouseDragNode(node[0]).Do(ctx)
	})
}

func mouseDragNode(n *cdp.Node) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		boxes, err := dom.GetContentQuads().WithNodeID(n.NodeID).Do(ctx)
		if err != nil {
			return err
		}

		box := boxes[0]
		c := len(box)
		if c%2 != 0 || c < 1 {
			return chromedp.ErrInvalidDimensions
		}

		var x, y float64
		for i := 0; i < c; i += 2 {
			x += box[i]
			y += box[i+1]
		}
		x /= float64(c / 2)
		y /= float64(c / 2)

		p := &input.DispatchMouseEventParams{
			Type:       input.MousePressed,
			X:          x,
			Y:          y,
			Button:     input.Left,
			ClickCount: 1,
		}

		if err := p.Do(ctx); err != nil {
			return err
		}

		// Mouse Move
		p.Type = input.MouseMoved
		p.X = x + 300

		if err := p.Do(ctx); err != nil {
			return err
		}

		p.Type = input.MouseReleased
		return p.Do(ctx)
	}
}
