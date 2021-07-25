package concurrency

type webschecker func(string) bool

func checkwebsite(wc webschecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		result[url] = wc(url)
	}

	return result

}
