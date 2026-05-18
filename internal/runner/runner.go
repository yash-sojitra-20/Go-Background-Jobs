package runner

func Run(task func()) {
	go task()
}