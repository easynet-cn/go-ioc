package ioc

type ContainerConfig struct {
	Name            string // Container struct name
	TemplateFile    string // Template file
	InputDirectory  string // Input directory
	OutputDirectory string // Output directory
	OutputFilename  string // Output go source filename
}
