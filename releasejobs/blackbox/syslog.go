package blackbox 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Syslog struct {

	/*Destination - Descr: transport protocol for syslog drain (udp/tcp/tls) Default: tls
*/
	Destination Destination `yaml:"destination,omitempty"`

	/*SourceDir - Descr: directory with subdirectories containing log files. log lines will be tagged with subdirectory name.
 Default: /var/vcap/sys/log
*/
	SourceDir interface{} `yaml:"source_dir,omitempty"`

}