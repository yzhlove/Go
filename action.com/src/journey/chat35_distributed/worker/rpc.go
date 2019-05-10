package worker

import "journey/chat31_love/engine"

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParseResult) error {
	var (
		engineReq engine.Request
		err       error
		engineRes engine.ParseResult
	)
	if engineReq, err = DeserializeRequest(req); err != nil {
		return err
	}
	if engineRes, err = engine.Worker(engineReq); err != nil {
		return err
	}
	*result = SerializeResult(engineRes)
	return nil
}
