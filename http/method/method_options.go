package method

import "github.com/litesoft-go/utils/options"

//noinspection GoUnusedExportedFunction
func AsOptions(src []Method, ifEmpty string) string {
	collector := options.For(len(src), ifEmpty)
	for _, v := range src {
		collector.Add(string(v))
	}
	return collector.Done()
}
