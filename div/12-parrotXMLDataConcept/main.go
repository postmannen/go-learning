/*
<project name="ardrone3" id="1">
	All ARDrone3-only commands
	<class name="Piloting" id="0">
		All commands related to piloting the drone
		<cmd name="FlatTrim" id="0">
			<comment
				title="Do a flat trim"
				desc="Do a flat trim of the accelerometer/gyro.\n
				Could be useful when the drone is sliding in hover mode."
				support="0901;090c;090e"
				result="Accelerometer and gyroscope are calibrated then event [FlatTrimChanged](#1-4-0) is triggered."/>
		</cmd>

		<cmd name="PCMD" id="2" buffer="NON_ACK">
			<comment
				title="Move the drone"
				desc="Move the drone.\n
				The libARController is sending the command each 50ms.\n\n
				**Please note that you should call setPilotingPCMD and not sendPilotingPCMD because the libARController is handling the periodicity and the buffer on which it is sent.**"
				support="0901;090c;090e"
				result="The drone moves! Yaaaaay!\n
				Event [SpeedChanged](#1-4-5), [AttitudeChanged](#1-4-6) and [PositionChanged](#1-4-4) (only if gps of the drone has fixed) are triggered."/>
			<arg name="flag" type="u8">
				Boolean flag: 1 if the roll and pitch values should be taken in consideration. 0 otherwise
			</arg>
			<arg name="roll" type="i8">
				Roll angle as signed percentage.
				On copters:
				Roll angle expressed as signed percentage of the max pitch/roll setting, in range [-100, 100]
				-100 corresponds to a roll angle of max pitch/roll to the left (drone will fly left)
				100 corresponds to a roll angle of max pitch/roll to the right (drone will fly right)
				This value may be clamped if necessary, in order to respect the maximum supported physical tilt of the copter.

				On fixed wings:
				Roll angle expressed as signed percentage of the physical max roll of the wing, in range [-100, 100]
				Negative value makes the plane fly to the left
				Positive value makes the plane fly to the right
			</arg>
*/

package main

import "fmt"

type Ardrone struct {
	Piloting
}

type Piloting struct {
}

func (p Piloting) FlatTrim() []byte {
	return []byte{0x01, 0x00, 0x00}
}

func (p Piloting) PCMD(flag uint8, roll int8, pitch int8, yaw int8, gaz int8, timestampAndSeqNum uint32) []byte {
	return []byte{0x01, 0x00, 0x02, flag, roll, pitch, yaw}
}

func main() {
	ardrone := Ardrone{}
	fmt.Printf("%#v\n", ardrone.Piloting.FlatTrim())
	fmt.Printf("%#v\n", ardrone.Piloting.PCMD())
}
