package polardb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBClusterSSL invokes the polardb.DescribeDBClusterSSL API synchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterssl.html
func (client *Client) DescribeDBClusterSSL(request *DescribeDBClusterSSLRequest) (response *DescribeDBClusterSSLResponse, err error) {
	response = CreateDescribeDBClusterSSLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterSSLWithChan invokes the polardb.DescribeDBClusterSSL API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterssl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterSSLWithChan(request *DescribeDBClusterSSLRequest) (<-chan *DescribeDBClusterSSLResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterSSL(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDBClusterSSLWithCallback invokes the polardb.DescribeDBClusterSSL API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterssl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterSSLWithCallback(request *DescribeDBClusterSSLRequest, callback func(response *DescribeDBClusterSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterSSLResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterSSL(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDBClusterSSLRequest is the request struct for api DescribeDBClusterSSL
type DescribeDBClusterSSLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterSSLResponse is the response struct for api DescribeDBClusterSSL
type DescribeDBClusterSSLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Items     []Item `json:"Items" xml:"Items"`
}

// CreateDescribeDBClusterSSLRequest creates a request to invoke DescribeDBClusterSSL API
func CreateDescribeDBClusterSSLRequest() (request *DescribeDBClusterSSLRequest) {
	request = &DescribeDBClusterSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterSSL", "polardb", "openAPI")
	return
}

// CreateDescribeDBClusterSSLResponse creates a response to parse from DescribeDBClusterSSL response
func CreateDescribeDBClusterSSLResponse() (response *DescribeDBClusterSSLResponse) {
	response = &DescribeDBClusterSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
