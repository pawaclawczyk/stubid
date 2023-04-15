package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	// todo: extract daemon routine
	if err := fasthttp.ListenAndServe(":8080", routeHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func xandrReadyHandler(ctx *fasthttp.RequestCtx) {
	// todo: extract response body to const
	ctx.Response.SetBody([]byte("1"))
}

func xandrBidderHandler(ctx *fasthttp.RequestCtx) {
	// todo: decode bid request
	// todo: respond with dummy bid response or dummy no bid response
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func xandrNotifyHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func xandrPixelHandler(ctx *fasthttp.RequestCtx) {
	// todo: respond with transparent pixel
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func xandrClickHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func xandrAuditNotifyHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func rtbNoticeWinHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func rtbNoticeLossHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func rtbNoticeBillingHandler(ctx *fasthttp.RequestCtx) {
	// todo: track request
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func routeHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/xandr/ready":
		xandrReadyHandler(ctx)
	case "/xandr/bid":
		xandrBidderHandler(ctx)
	case "/xandr/notify":
		xandrNotifyHandler(ctx)
	case "/xandr/pixel":
		xandrPixelHandler(ctx)
	case "/xandr/click":
		xandrClickHandler(ctx)
	case "/xandr/audit-notify":
		xandrAuditNotifyHandler(ctx)
	case "/rtb/notice/win":
		rtbNoticeWinHandler(ctx)
	case "/rtb/notice/loss":
		rtbNoticeLossHandler(ctx)
	case "/rtb/notice/billing":
		rtbNoticeBillingHandler(ctx)
	default:
		log.Printf("error in routeHandler: unknown path: %s", ctx.Path())
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}
