# StuBid: Stupid Bidder

## Test Xandr endpoints

```shell
curl -i localhost:8080/xandr/ready
```

```shell
curl -i -X POST -d @samples/xandr-bid-request-banner.json localhost:8080/xandr/bid
```

```shell
curl -i -X POST -H "stubid-no-bid: 1" -d @samples/xandr-bid-request-banner.json localhost:8080/xandr/bid
```
