package cmd

func Execute(VERSION, COMMIT string) {
	version = VERSION
	versionCommit = COMMIT
	if err := rootCmd.Execute(); err != nil {
		klog.Exit(err)
	}
}
