// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"define"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetOrder(string order_id, bool sub_order)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetSubOrder(i64 id)")
	fmt.Fprintln(os.Stderr, "  ComplexOrder GetSubOrderByNo(string orderNo)")
	fmt.Fprintln(os.Stderr, "   GetSubOrderItems(i64 subOrderId)")
	fmt.Fprintln(os.Stderr, "  Result64 SubmitTradeOrder(ComplexOrder o, double rate)")
	fmt.Fprintln(os.Stderr, "  Result64 TradeOrderCashPay(i64 orderId)")
	fmt.Fprintln(os.Stderr, "  Result64 TradeOrderUpdateTicket(i64 orderId, string img)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := define.NewSaleServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "GetOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetOrder requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2) == "true"
		value1 := argvalue1
		fmt.Print(client.GetOrder(value0, value1))
		fmt.Print("\n")
		break
	case "GetSubOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrder requires 1 args")
			flag.Usage()
		}
		argvalue0, err195 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err195 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetSubOrder(value0))
		fmt.Print("\n")
		break
	case "GetSubOrderByNo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrderByNo requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetSubOrderByNo(value0))
		fmt.Print("\n")
		break
	case "GetSubOrderItems":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSubOrderItems requires 1 args")
			flag.Usage()
		}
		argvalue0, err197 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err197 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetSubOrderItems(value0))
		fmt.Print("\n")
		break
	case "SubmitTradeOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SubmitTradeOrder requires 2 args")
			flag.Usage()
		}
		arg198 := flag.Arg(1)
		mbTrans199 := thrift.NewTMemoryBufferLen(len(arg198))
		defer mbTrans199.Close()
		_, err200 := mbTrans199.WriteString(arg198)
		if err200 != nil {
			Usage()
			return
		}
		factory201 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt202 := factory201.GetProtocol(mbTrans199)
		argvalue0 := define.NewComplexOrder()
		err203 := argvalue0.Read(jsProt202)
		if err203 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err204 := (strconv.ParseFloat(flag.Arg(2), 64))
		if err204 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.SubmitTradeOrder(value0, value1))
		fmt.Print("\n")
		break
	case "TradeOrderCashPay":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TradeOrderCashPay requires 1 args")
			flag.Usage()
		}
		argvalue0, err205 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err205 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TradeOrderCashPay(value0))
		fmt.Print("\n")
		break
	case "TradeOrderUpdateTicket":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TradeOrderUpdateTicket requires 2 args")
			flag.Usage()
		}
		argvalue0, err206 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err206 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.TradeOrderUpdateTicket(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}