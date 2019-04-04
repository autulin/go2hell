package jobs

type HelloJob struct{}

func (h *HelloJob) Init() {
	println("here do something init the job.")
}

func (h *HelloJob) Exe() {
	println("say hello!")
}

func (h *HelloJob) End() {
	println("here do something when job finished.")
}
