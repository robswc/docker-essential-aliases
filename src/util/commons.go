package util

const WorkingContainerId = "WORKING_CONTAINER_ID"
const WorkingContainerName = "WORKING_CONTAINER_NAME"
const TempFileName = "dea-temp.yml"
const TempFileLocation = "/tmp"
const DeaVersion = "0.0.1"

func GetTempFile() string {
	return TempFileLocation + "/" + TempFileName
}
