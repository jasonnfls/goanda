package goanda

const (
	streamUrl = "https://stream-fxtrade.oanda.com/v1"
	apiUrl    = "https://api-fxtrade.oanda.com/v1"

	// may need to change these to uppercase.. check this
)

var (
	CurrentAccount int
	Accounts       = make(map[int]string)
	Instruments    []string
	SessionIds     [2]string
)

/*
Copyright (c) 2016 Samuel B. Sendelbach

Permission is hereby granted, free of charge,
to any person obtaining a copy of this software
and associated documentation files (the "Software"),
to deal in the Software without restriction,
including without limitation the rights to
use, copy, modify,
merge, publish, distribute,
sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

NOTE THIS IS NOT SOFTWARE OFFICIAL SUPPORTED BY OANDA, NOR IS THIS SOFTWARE AFFILIATED
WITH THE OANDA CORPORATION.

Leverage trading is high risk and not suitable for all.
You could lose all of your deposited funds.
*/
