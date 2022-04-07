// Code generated by protoc-gen-volcengine-sdk
// source: VodMediaService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/request"
)

func Test_UpdateMediaInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaInfoRequest{
		Vid:         "your Vid",
		PosterUri:   nil,
		Title:       nil,
		Description: nil,
		Tags:        nil,
	}

	resp, status, err := instance.UpdateMediaInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateMediaPublishStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaPublishStatusRequest{
		Vid:    "your Vid",
		Status: "your Status",
	}

	resp, status, err := instance.UpdateMediaPublishStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaInfos(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaInfosRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetMediaInfos(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetRecommendedPoster(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetRecommendedPosterRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetRecommendedPoster(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteMedia(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteMediaRequest{
		Vids:         "your Vids",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteMedia(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteTranscodes(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteTranscodesRequest{
		Vid:          "your Vid",
		FileIds:      "your FileIds",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteTranscodes(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaListRequest{
		SpaceName: "your SpaceName",
		Vid:       "your Vid",
		Status:    "your Status",
		Order:     "your Order",
		Tags:      "your Tags",
		StartTime: "your StartTime",
		EndTime:   "your EndTime",
		Offset:    "your Offset",
		PageSize:  "your PageSize",
	}

	resp, status, err := instance.GetMediaList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetSubtitleInfoList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetSubtitleInfoListRequest{
		Vid:         "your Vid",
		FileIds:     "your FileIds",
		Languages:   "your Languages",
		Formats:     "your Formats",
		LanguageIds: "your LanguageIds",
		SubtitleIds: "your SubtitleIds",
		Status:      "your Status",
		Title:       "your Title",
		Tag:         "your Tag",
		Offset:      "your Offset",
		PageSize:    "your PageSize",
		Ssl:         "your Ssl",
	}

	resp, status, err := instance.GetSubtitleInfoList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleStatusRequest{
		Vid:       "your Vid",
		FileIds:   "your FileIds",
		Languages: "your Languages",
		Formats:   "your Formats",
		Status:    "your Status",
	}

	resp, status, err := instance.UpdateSubtitleStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleInfoRequest{
		Vid:      "your Vid",
		FileId:   "your FileId",
		Language: "your Language",
		Format:   "your Format",
		Title:    nil,
		Tag:      nil,
	}

	resp, status, err := instance.UpdateSubtitleInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
