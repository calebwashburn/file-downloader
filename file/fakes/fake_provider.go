// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/file-downloader-resource/file"
)

type FakeProvider struct {
	DownloadFileStub        func(targetDirectory, productSlug, version, pattern string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		targetDirectory string
		productSlug     string
		version         string
		pattern         string
	}
	downloadFileReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProvider) DownloadFile(targetDirectory string, productSlug string, version string, pattern string) error {
	fake.downloadFileMutex.Lock()
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		targetDirectory string
		productSlug     string
		version         string
		pattern         string
	}{targetDirectory, productSlug, version, pattern})
	fake.recordInvocation("DownloadFile", []interface{}{targetDirectory, productSlug, version, pattern})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(targetDirectory, productSlug, version, pattern)
	} else {
		return fake.downloadFileReturns.result1
	}
}

func (fake *FakeProvider) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeProvider) DownloadFileArgsForCall(i int) (string, string, string, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.downloadFileArgsForCall[i].targetDirectory, fake.downloadFileArgsForCall[i].productSlug, fake.downloadFileArgsForCall[i].version, fake.downloadFileArgsForCall[i].pattern
}

func (fake *FakeProvider) DownloadFileReturns(result1 error) {
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ file.Provider = new(FakeProvider)