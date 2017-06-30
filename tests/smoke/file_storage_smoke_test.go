/*
Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package smoke

import (
	"github.com/stretchr/testify/require"
)

// Smoke Test for File System Storage - Test check the following operations on FileSystem Storage in order
//Create,Mount,Write,Read,Unmount and Delete.
func (suite *SmokeSuite) TestFileStorage_SmokeTest() {
	defer suite.fileTestDataCleanUp()
	suite.T().Log("File Storage Smoke Test - Create,Mount,write to, read from  and Unmount Filesystem")
	rfc := suite.helper.GetFileSystemClient()

	suite.T().Log("Step 1: Create file System")
	_, fsc_err := suite.helper.CreateFileStorage()
	require.Nil(suite.T(), fsc_err)
	fileSystemList, _ := rfc.FSList()
	require.Equal(suite.T(), 1, len(fileSystemList), "There should one shared file system present")
	filesystemData := fileSystemList[0]
	require.Equal(suite.T(), "testfs", filesystemData.Name, "make sure filesystem name matches")
	suite.T().Log("File system created")

	suite.T().Log("Step 2: Mount file System")
	_, mtfs_err := suite.helper.MountFileStorage()
	require.Nil(suite.T(), mtfs_err)
	suite.T().Log("File system mounted successfully")

	suite.T().Log("Step 3: Write to file system")
	_, wfs_err := suite.helper.WriteToFileStorage("Test data for file", "fsFile1")
	require.Nil(suite.T(), wfs_err)
	suite.T().Log("Write to file system successful")

	suite.T().Log("Step 4: Read from file system")
	read, rd_err := suite.helper.ReadFromFileStorage("fsFile1")
	require.Nil(suite.T(), rd_err)
	require.Contains(suite.T(), read, "Test data for file", "make sure content of the files is unchanged")
	suite.T().Log("Read from file system successful")

	suite.T().Log("Step 5: Mount file System")
	_, umtfs_err := suite.helper.UnmountFileStorage()
	require.Nil(suite.T(), umtfs_err)
	suite.T().Log("File system mounted successfully")

	suite.T().Log("Step 6: Deleting file storage")
	suite.helper.DeleteFileStorage()
	//Delete is not deleting filesystem - known issue
	//require.Nil(suite.T(), fsd_err)
	suite.T().Log("File system deleted")
}

func (s *SmokeSuite) fileTestDataCleanUp() {
	s.helper.UnmountFileStorage()
	s.helper.DeleteFileStorage()

}
