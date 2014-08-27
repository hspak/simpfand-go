package main

const FAN_PATH string = "/proc/acpi/ibm/fan"

const TEMP_PATH_CNT int = 2
const TEMP_PATH_1 string = "/sys/devices/platform/coretemp.0/temp1_input"
const TEMP_PATH_2 string = "/sys/devices/platform/coretemp.0/hwmon/hwmon1/temp1_input"

const CONFIG_PATH string = "/etc/simpfand.conf"
