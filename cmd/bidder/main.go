package main

import (
	"encoding/json"
	"github.com/prebid/openrtb/v19/openrtb2"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
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
	req, err := fastjson.ParseBytes(ctx.PostBody())
	if err != nil {
		log.Printf("xandrBidderHandler: error in parsing bid request: %v", err)
		ctx.Error("internal error", fasthttp.StatusInternalServerError)
		return
	}
	if len(ctx.Request.Header.Peek("stubid-no-bid")) > 0 {
		log.Println("xandrBidderHandler: testing: enforcing no bid response")
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}
	// todo: replace static dummy bid response with internal auction and bid response builder
	res := openrtb2.BidResponse{
		ID: string(req.GetStringBytes("id")),
		SeatBid: []openrtb2.SeatBid{
			{
				Bid: []openrtb2.Bid{
					{
						ID:    "bid-123",
						ImpID: string(req.GetStringBytes("imp", "0", "id")),
						Price: 1.23,
						AdID:  "xandr-creative-id-123",
						NURL:  "localhost:8080/rtb/notice/win",
						LURL:  "localhost:8080/rtb/notice/loss",
						CID:   "campaign-id-123",
					},
				},
				Seat: "xandr-member-id-123",
			},
		},
		BidID: "bid-response-123",
		Cur:   "USD",
	}
	bRes, err := json.Marshal(res)
	if err != nil {
		log.Printf("xandrBidderHandler: error in encoding bid response: %v", err)
		ctx.Error("internal error", fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetBody(bRes)
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
