package main

func main() {
	Execute()
}
func init() {
	rootCmd.AddCommand(grpcCmd)
	rootCmd.AddCommand(progressCmd)
	rootCmd.AddCommand(notifySubscribeCmd)
	rootCmd.AddCommand(RegisterCmd)
	rootCmd.AddCommand(syncCmd)
	mangasCmd.AddCommand(mangaSearcshCmd)
	rootCmd.AddCommand(mangasCmd)
	rootCmd.AddCommand(libraryCmd)
	libraryCmd.AddCommand(libraryAddCmd)
	rootCmd.AddCommand(addprogressCmd)
	progressCmd.AddCommand(addprogressUpdateCmd)
	mangaCmd.AddCommand(mangaSearchCmd)

}
