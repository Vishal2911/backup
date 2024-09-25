package util

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func SetLoger() *logrus.Logger {

	// Define a command-line flag for log level
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "Log level (debug, info, warn, error)")

	// Parse the command-line flags
	flag.Parse()
	// Set the log level based on the flag value
	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
	return Logger
}

func Log(logLevel, pkgLevel, functionName string, message, parameater interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameater == nil {
			Logger.Debugf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Debugf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	case model.LogLevelWarning:
		if parameater == nil {
			Logger.Warnf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Warnf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	case model.LogLevelError:
		if parameater == nil {
			Logger.Errorf("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Errorf("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.LogLevelError, parameater)
		}
	default:
		if parameater == nil {
			Logger.Infof("%s, %s, %v\n", pkgLevel, functionName, message)
		} else {
			Logger.Infof("%s, %s, %v, %s = %v\n", pkgLevel, functionName, message, model.Value, parameater)
		}
	}
}

// Client represents interface for Kite Connect client.
type Client struct {
	clientCode  string
	password    string
	accessToken string
	debug       bool
	baseURI     string
	apiKey      string
	httpClient  HTTPClient
}

func getLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("please check your network connection")
}

func getIpAndMac() (string, string, string, error) {

	//----------------------
	// Get the local machine IP address
	//----------------------

	var localIp, currentNetworkHardwareName string

	localIp, err := getLocalIP()

	if err != nil {
		return "", "", "", err
	}

	// get all the system's or local machine's network interfaces

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), localIp) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		return "", "", "", err
	}

	macAddress := netInterface.HardwareAddr

	// verify if the MAC address can be parsed properly
	_, err = net.ParseMAC(macAddress.String())

	if err != nil {
		return "", "", "", err
	}

	publicIp, err := getPublicIp()
	if err != nil {
		return "", "", "", err
	}

	return localIp, publicIp, macAddress.String(), nil

}

func getPublicIp() (string, error) {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return "", err
	}

	content, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (c *Client) doEnvelope(method, uri string, params map[string]interface{}, headers http.Header, v interface{}, authorization ...bool) error {
	if params == nil {
		params = map[string]interface{}{}
	}

	// Send custom headers set
	if headers == nil {
		headers = map[string][]string{}
	}

	localIp, publicIp, mac, err := getIpAndMac()

	if err != nil {
		return err
	}

	// Add Kite Connect version to header
	headers.Add("Content-Type", "application/json")
	headers.Add("X-ClientLocalIP", localIp)
	headers.Add("X-ClientPublicIP", publicIp)
	headers.Add("X-MACAddress", mac)
	headers.Add("Accept", "application/json")
	headers.Add("X-UserType", "USER")
	headers.Add("X-SourceID", "WEB")
	headers.Add("X-PrivateKey", c.apiKey)
	if authorization != nil && authorization[0] {
		headers.Add("Authorization", "Bearer "+c.accessToken)
	}

	return c.httpClient.DoEnvelope(method, c.baseURI+uri, params, headers, v)
}

type HTTPClient interface {
	Do(method, rURL string, params map[string]interface{}, headers http.Header) (HTTPResponse, error)
	DoEnvelope(method, url string, params map[string]interface{}, headers http.Header, obj interface{}) error
	GetClient() *httpClient
}

type httpClient struct {
	client *http.Client
	hLog   *log.Logger
	debug  bool
}

// HTTPResponse encompasses byte body  + the response of an HTTP request.
type HTTPResponse struct {
	Body     []byte
	Response *http.Response
}