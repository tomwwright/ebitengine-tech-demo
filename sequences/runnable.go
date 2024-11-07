package sequences

type Done func()
type Runnable interface {
	Run(done Done)
}
