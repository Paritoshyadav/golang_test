package concurrency

type webschecker func(string) bool

type res struct {
	url    string
	status bool
}

func checkwebsite(wc webschecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	resultchannel := make(chan res)

	for _, url := range urls {
		go func(u string) {

			resultchannel <- res{u, wc(u)}

		}(url)

	}

	for i := 0; i < len(urls); i++ {
		u := <-resultchannel
		result[u.url] = u.status

	}

	return result

}
