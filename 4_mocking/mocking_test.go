package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	mocks "avalith/testing/mocking/mocks"
)

// mockgen -destination=mocks/mock_IDManager.go -package=mocks avalith/testing/mocking/interfaces IDManager
func TestMain(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Init vars
	var (
		IDGenMock = mocks.NewMockIDManager(mockCtrl)
		itemName  = "Name"
		genID     = "001"
	)

	t.Run("NewItem", func(t *testing.T) {

		t.Run("SHOULD return created item", func(t *testing.T) {
			gomock.InOrder(
				IDGenMock.EXPECT().GenID().Return(genID).Times(1),
			)

			res := NewItem(IDGenMock, itemName)
			assert.Assert(t, res.Name == itemName, fmt.Sprintf("%s not equal %s", res.Name, itemName))
			assert.Assert(t, res.ID == genID, fmt.Sprintf("%s not equal %s", res.ID, genID))
		})

	})
}
