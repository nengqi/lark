// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Attendance_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Attendance()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserSettings(ctx, &lark.UpdateUserSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateGroup(ctx, &lark.CreateUpdateGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteGroup(ctx, &lark.DeleteGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetGroup(ctx, &lark.GetGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateShift(ctx, &lark.CreateShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteShift(ctx, &lark.DeleteShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetShiftByID(ctx, &lark.GetShiftByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetShiftByName(ctx, &lark.GetShiftByNameReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetStatisticsData(ctx, &lark.GetStatisticsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetStatisticsHeader(ctx, &lark.GetStatisticsHeaderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserStatisticsSettings(ctx, &lark.UpdateUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserStatisticsSettings(ctx, &lark.GetUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserDailyShift(ctx, &lark.GetUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserTask(ctx, &lark.GetUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserFlow(ctx, &lark.GetUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUserFlow(ctx, &lark.BatchGetUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateUserFlow(ctx, &lark.BatchCreateUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserTaskRemedy(ctx, &lark.GetUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateUserDailyShift(ctx, &lark.CreateUpdateUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserApproval(ctx, &lark.GetUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUserApproval(ctx, &lark.CreateUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Attendance()

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUpdateUserSettings(func(ctx context.Context, request *lark.UpdateUserSettingsReq, options ...lark.MethodOptionFunc) (*lark.UpdateUserSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateUserSettings()

			_, _, err := moduleCli.UpdateUserSettings(ctx, &lark.UpdateUserSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUploadAttendanceFile(func(ctx context.Context, request *lark.UploadAttendanceFileReq, options ...lark.MethodOptionFunc) (*lark.UploadAttendanceFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUploadAttendanceFile()

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateUpdateGroup(func(ctx context.Context, request *lark.CreateUpdateGroupReq, options ...lark.MethodOptionFunc) (*lark.CreateUpdateGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateUpdateGroup()

			_, _, err := moduleCli.CreateUpdateGroup(ctx, &lark.CreateUpdateGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceDeleteGroup(func(ctx context.Context, request *lark.DeleteGroupReq, options ...lark.MethodOptionFunc) (*lark.DeleteGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteGroup()

			_, _, err := moduleCli.DeleteGroup(ctx, &lark.DeleteGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetGroup(func(ctx context.Context, request *lark.GetGroupReq, options ...lark.MethodOptionFunc) (*lark.GetGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetGroup()

			_, _, err := moduleCli.GetGroup(ctx, &lark.GetGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateShift(func(ctx context.Context, request *lark.CreateShiftReq, options ...lark.MethodOptionFunc) (*lark.CreateShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateShift()

			_, _, err := moduleCli.CreateShift(ctx, &lark.CreateShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceDeleteShift(func(ctx context.Context, request *lark.DeleteShiftReq, options ...lark.MethodOptionFunc) (*lark.DeleteShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteShift()

			_, _, err := moduleCli.DeleteShift(ctx, &lark.DeleteShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetShiftByID(func(ctx context.Context, request *lark.GetShiftByIDReq, options ...lark.MethodOptionFunc) (*lark.GetShiftByIDResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetShiftByID()

			_, _, err := moduleCli.GetShiftByID(ctx, &lark.GetShiftByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetShiftByName(func(ctx context.Context, request *lark.GetShiftByNameReq, options ...lark.MethodOptionFunc) (*lark.GetShiftByNameResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetShiftByName()

			_, _, err := moduleCli.GetShiftByName(ctx, &lark.GetShiftByNameReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetStatisticsData(func(ctx context.Context, request *lark.GetStatisticsDataReq, options ...lark.MethodOptionFunc) (*lark.GetStatisticsDataResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetStatisticsData()

			_, _, err := moduleCli.GetStatisticsData(ctx, &lark.GetStatisticsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetStatisticsHeader(func(ctx context.Context, request *lark.GetStatisticsHeaderReq, options ...lark.MethodOptionFunc) (*lark.GetStatisticsHeaderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetStatisticsHeader()

			_, _, err := moduleCli.GetStatisticsHeader(ctx, &lark.GetStatisticsHeaderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUpdateUserStatisticsSettings(func(ctx context.Context, request *lark.UpdateUserStatisticsSettingsReq, options ...lark.MethodOptionFunc) (*lark.UpdateUserStatisticsSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateUserStatisticsSettings()

			_, _, err := moduleCli.UpdateUserStatisticsSettings(ctx, &lark.UpdateUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserStatisticsSettings(func(ctx context.Context, request *lark.GetUserStatisticsSettingsReq, options ...lark.MethodOptionFunc) (*lark.GetUserStatisticsSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserStatisticsSettings()

			_, _, err := moduleCli.GetUserStatisticsSettings(ctx, &lark.GetUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserDailyShift(func(ctx context.Context, request *lark.GetUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.GetUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserDailyShift()

			_, _, err := moduleCli.GetUserDailyShift(ctx, &lark.GetUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserTask(func(ctx context.Context, request *lark.GetUserTaskReq, options ...lark.MethodOptionFunc) (*lark.GetUserTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserTask()

			_, _, err := moduleCli.GetUserTask(ctx, &lark.GetUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserFlow(func(ctx context.Context, request *lark.GetUserFlowReq, options ...lark.MethodOptionFunc) (*lark.GetUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserFlow()

			_, _, err := moduleCli.GetUserFlow(ctx, &lark.GetUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceBatchGetUserFlow(func(ctx context.Context, request *lark.BatchGetUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchGetUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchGetUserFlow()

			_, _, err := moduleCli.BatchGetUserFlow(ctx, &lark.BatchGetUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceBatchCreateUserFlow(func(ctx context.Context, request *lark.BatchCreateUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchCreateUserFlow()

			_, _, err := moduleCli.BatchCreateUserFlow(ctx, &lark.BatchCreateUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserTaskRemedy(func(ctx context.Context, request *lark.GetUserTaskRemedyReq, options ...lark.MethodOptionFunc) (*lark.GetUserTaskRemedyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserTaskRemedy()

			_, _, err := moduleCli.GetUserTaskRemedy(ctx, &lark.GetUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateUpdateUserDailyShift(func(ctx context.Context, request *lark.CreateUpdateUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.CreateUpdateUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateUpdateUserDailyShift()

			_, _, err := moduleCli.CreateUpdateUserDailyShift(ctx, &lark.CreateUpdateUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetUserApproval(func(ctx context.Context, request *lark.GetUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.GetUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetUserApproval()

			_, _, err := moduleCli.GetUserApproval(ctx, &lark.GetUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateUserApproval(func(ctx context.Context, request *lark.CreateUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.CreateUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateUserApproval()

			_, _, err := moduleCli.CreateUserApproval(ctx, &lark.CreateUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Attendance()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserSettings(ctx, &lark.UpdateUserSettingsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateGroup(ctx, &lark.CreateUpdateGroupReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteGroup(ctx, &lark.DeleteGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetGroup(ctx, &lark.GetGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateShift(ctx, &lark.CreateShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteShift(ctx, &lark.DeleteShiftReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetShiftByID(ctx, &lark.GetShiftByIDReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetShiftByName(ctx, &lark.GetShiftByNameReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetStatisticsData(ctx, &lark.GetStatisticsDataReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetStatisticsHeader(ctx, &lark.GetStatisticsHeaderReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateUserStatisticsSettings(ctx, &lark.UpdateUserStatisticsSettingsReq{
				UserStatsViewID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserStatisticsSettings(ctx, &lark.GetUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserDailyShift(ctx, &lark.GetUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserTask(ctx, &lark.GetUserTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserFlow(ctx, &lark.GetUserFlowReq{
				UserFlowID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetUserFlow(ctx, &lark.BatchGetUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateUserFlow(ctx, &lark.BatchCreateUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserTaskRemedy(ctx, &lark.GetUserTaskRemedyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateUserDailyShift(ctx, &lark.CreateUpdateUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetUserApproval(ctx, &lark.GetUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUserApproval(ctx, &lark.CreateUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
