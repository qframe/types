package qtypes_interface

type QPlugin interface {
	// Starts the plugin and does not quit until a fatal error occurs
	Run() error
}