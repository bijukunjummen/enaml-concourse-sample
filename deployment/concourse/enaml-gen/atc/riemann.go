package atc 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Riemann struct {

	/*Host - Descr: If configured, detailed metrics will be emitted to the specified Riemann
server.
 Default: 
*/
	Host interface{} `yaml:"host,omitempty"`

	/*Port - Descr: Port of the Riemann server to emit events to.
 Default: 5555
*/
	Port interface{} `yaml:"port,omitempty"`

}