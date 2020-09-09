package main

import (
	"fmt"
)

type projectDef uint8 
type classDef uint8
type cmdDef uint16

type command struct {
	project projectDef
	class   classDef
	cmd     cmdDef
}

// All ARDrone3-only commands
	const ardrone3 projectDef = 1
// All commands related to piloting the drone
const piloting classDef = 0
// title : Take off, 
// desc : Ask the drone to take off.\n On the fixed wings (such as Disco): not used except to cancel a land., 
// support : 0901;090c;090e, 
// result : On the quadcopters: the drone takes off if its [FlyingState](#1-4-1) was landed.\n On the fixed wings, the landing process is aborted if the [FlyingState](#1-4-1) was landing.\n Then, event [FlyingState](#1-4-1) is triggered., 
const takeOff cmdDef = 1

type ardrone3PilotingTakeOff command

func (a ardrone3PilotingTakeOff) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingTakeOff = ardrone3PilotingTakeOff {
project: ardrone3,
class: piloting,
cmd: takeOff,
}

// title : Move the drone, 
// desc : Move the drone.\n The libARController is sending the command each 50ms.\n\n **Please note that you should call setPilotingPCMD and not sendPilotingPCMD because the libARController is handling the periodicity and the buffer on which it is sent.**, 
// support : 0901;090c;090e, 
// result : The drone moves! Yaaaaay!\n Event [SpeedChanged](#1-4-5), [AttitudeChanged](#1-4-6) and [PositionChanged](#1-4-4) (only if gps of the drone has fixed) are triggered., 
const pCMD cmdDef = 2

type ardrone3PilotingPCMD command

func (a ardrone3PilotingPCMD) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingPCMD = ardrone3PilotingPCMD {
project: ardrone3,
class: piloting,
cmd: pCMD,
}

// title : Land, 
// desc : Land.\n Please note that on copters, if you put some positive gaz (in the [PilotingCommand](#1-0-2)) during the landing, it will cancel it., 
// support : 0901;090c;090e, 
// result : On the copters, the drone lands if its [FlyingState](#1-4-1) was taking off, hovering or flying.\n On the fixed wings, the drone lands if its [FlyingState](#1-4-1) was hovering or flying.\n Then, event [FlyingState](#1-4-1) is triggered., 
const landing cmdDef = 3

type ardrone3PilotingLanding command

func (a ardrone3PilotingLanding) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingLanding = ardrone3PilotingLanding {
project: ardrone3,
class: piloting,
cmd: landing,
}

// title : Cut out the motors, 
// desc : Cut out the motors.\n This cuts immediatly the motors. The drone will fall.\n This command is sent on a dedicated high priority buffer which will infinitely retry to send it if the command is not delivered., 
// support : 0901;090c;090e, 
// result : The drone immediatly cuts off its motors.\n Then, event [FlyingState](#1-4-1) is triggered., 
const emergency cmdDef = 4

type ardrone3PilotingEmergency command

func (a ardrone3PilotingEmergency) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingEmergency = ardrone3PilotingEmergency {
project: ardrone3,
class: piloting,
cmd: emergency,
}

// title : Return home, 
// desc : Return home.\n Ask the drone to fly to its [HomePosition](#1-24-0).\n The availability of the return home can be get from [ReturnHomeState](#1-4-3).\n Please note that the drone will wait to be hovering to start its return home. This means that it will wait to have a [flag](#1-0-2) set at 0., 
// support : 0901;090c;090e, 
// result : The drone will fly back to its home position.\n Then, event [ReturnHomeState](#1-4-3) is triggered.\n You can get a state pending if the drone is not ready to start its return home process but will do it as soon as it is possible., 
const navigateHome cmdDef = 5

type ardrone3PilotingNavigateHome command

func (a ardrone3PilotingNavigateHome) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingNavigateHome = ardrone3PilotingNavigateHome {
project: ardrone3,
class: piloting,
cmd: navigateHome,
}

// title : Auto take off mode, 
// desc : Auto take off mode., 
const autoTakeOffMode cmdDef = 6

type ardrone3PilotingAutoTakeOffMode command

func (a ardrone3PilotingAutoTakeOffMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingAutoTakeOffMode = ardrone3PilotingAutoTakeOffMode {
project: ardrone3,
class: piloting,
cmd: autoTakeOffMode,
}

// title : Move the drone to a relative position, 
// desc : Move the drone to a relative position and rotate heading by a given angle.\n Moves are relative to the current drone orientation, (drone's reference).\n Also note that the given rotation will not modify the move (i.e. moves are always rectilinear)., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The drone will move of the given offsets.\n Then, event [RelativeMoveEnded](#1-34-0) is triggered.\n If you send a second relative move command, the drone will trigger a [RelativeMoveEnded](#1-34-0) with the offsets it managed to do before this new command and the value of error set to interrupted., 
const moveBy cmdDef = 7

type ardrone3PilotingmoveBy command

func (a ardrone3PilotingmoveBy) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingmoveBy = ardrone3PilotingmoveBy {
project: ardrone3,
class: piloting,
cmd: moveBy,
}

// title : Prepare the drone to take off, 
// desc : Prepare the drone to take off.\n On copters: initiates the thrown takeoff. Note that the drone will do the thrown take off even if it is steady.\n On fixed wings: initiates the take off process on the fixed wings.\n\n Setting the state to 0 will cancel the preparation. You can cancel it before that the drone takes off., 
// support : 090e;090c:4.3.0, 
// result : The drone will arm its motors if not already armed.\n Then, event [FlyingState](#1-4-1) is triggered with state set at motor ramping.\n Then, event [FlyingState](#1-4-1) is triggered with state set at userTakeOff.\n Then user can throw the drone to make it take off., 
const userTakeOff cmdDef = 8

type ardrone3PilotingUserTakeOff command

func (a ardrone3PilotingUserTakeOff) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingUserTakeOff = ardrone3PilotingUserTakeOff {
project: ardrone3,
class: piloting,
cmd: userTakeOff,
}

// title : Circle, 
// desc : Make the fixed wing circle.\n The circle will use the [CirclingAltitude](#1-6-14) and the [CirclingRadius](#1-6-13), 
// support : 090e, 
// result : The fixed wing will circle in the given direction.\n Then, event [FlyingState](#1-4-1) is triggered with state set at hovering., 
const circle cmdDef = 9

type ardrone3PilotingCircle command

func (a ardrone3PilotingCircle) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingCircle = ardrone3PilotingCircle {
project: ardrone3,
class: piloting,
cmd: circle,
}

// title : Move to a location, 
// desc : Move the drone to a specified location.\n If a new command moveTo is sent, the drone will immediatly run it (no cancel will be issued).\n If a [CancelMoveTo](#1-0-11) command is sent, the moveTo is stopped.\n During the moveTo, all pitch, roll and gaz values of the piloting command will be ignored by the drone.\n However, the yaw value can be used., 
// support : 090c:4.3.0, 
// result : Event [MovingTo](#1-4-12) is triggered with state running. Then, the drone will move to the given location.\n Then, event [MoveToChanged](#1-4-12) is triggered with state succeed., 
const moveTo cmdDef = 10

type ardrone3PilotingmoveTo command

func (a ardrone3PilotingmoveTo) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingmoveTo = ardrone3PilotingmoveTo {
project: ardrone3,
class: piloting,
cmd: moveTo,
}

// title : Cancel the moveTo, 
// desc : Cancel the current moveTo.\n If there is no current moveTo, this command has no effect., 
// support : 090c:4.3.0, 
// result : Event [MoveToChanged](#1-4-12) is triggered with state canceled., 
const cancelMoveTo cmdDef = 11

type ardrone3PilotingCancelMoveTo command

func (a ardrone3PilotingCancelMoveTo) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingCancelMoveTo = ardrone3PilotingCancelMoveTo {
project: ardrone3,
class: piloting,
cmd: cancelMoveTo,
}

// title : Start a piloted POI, 
// desc : Start a piloted Point Of Interest.\n During a piloted POI, the drone will always look at the given POI but can be piloted normally. However, yaw value is ignored. Camera tilt and pan command is also ignored.\n Ignored if [PilotedPOI](#1-4-14) state is UNAVAILABLE., 
// support : 090c:4.3.0, 
// result : If the drone is hovering, event [PilotedPOI](#1-4-14) is triggered with state RUNNING. If the drone is not hovering, event [PilotedPOI](#1-4-14) is triggered with state PENDING, waiting to hover. When the drone hovers, the state will change to RUNNING. If the drone does not hover for a given time, piloted POI is canceled by the drone and state will change to AVAILABLE. Then, the drone will look at the given location., 
const startPilotedPOI cmdDef = 12

type ardrone3PilotingStartPilotedPOI command

func (a ardrone3PilotingStartPilotedPOI) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStartPilotedPOI = ardrone3PilotingStartPilotedPOI {
project: ardrone3,
class: piloting,
cmd: startPilotedPOI,
}

// title : Stop the piloted POI, 
// desc : Stop the piloted Point Of Interest.\n If [PilotedPOI](#1-4-14) state is RUNNING or PENDING, stop it., 
// support : 090c:4.3.0, 
// result : Event [PilotedPOI](#1-4-14) is triggered with state AVAILABLE., 
const stopPilotedPOI cmdDef = 13

type ardrone3PilotingStopPilotedPOI command

func (a ardrone3PilotingStopPilotedPOI) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStopPilotedPOI = ardrone3PilotingStopPilotedPOI {
project: ardrone3,
class: piloting,
cmd: stopPilotedPOI,
}

// title : Cancel the relative move, 
// desc : Cancel the current relative move.\n If there is no current relative move, this command has no effect., 
// result : Event [RelativeMoveChanged](#1-4-16) is triggered with state canceled., 
const cancelMoveBy cmdDef = 14

type ardrone3PilotingCancelMoveBy command

func (a ardrone3PilotingCancelMoveBy) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingCancelMoveBy = ardrone3PilotingCancelMoveBy {
project: ardrone3,
class: piloting,
cmd: cancelMoveBy,
}

// Animation commands
const animations classDef = 5
// title : Make a flip, 
// desc : Make a flip., 
// support : 0901;090c, 
// result : The drone will make a flip if it has enough battery., 
const flip cmdDef = 0

type ardrone3AnimationsFlip command

func (a ardrone3AnimationsFlip) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var animationsFlip = ardrone3AnimationsFlip {
project: ardrone3,
class: animations,
cmd: flip,
}

// Ask the drone to move camera
const camera classDef = 1
// title : Move the camera, 
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientation](#1-25-0) is triggered., 
const orientation cmdDef = 0

type ardrone3CameraOrientation command

func (a ardrone3CameraOrientation) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraOrientation = ardrone3CameraOrientation {
project: ardrone3,
class: camera,
cmd: orientation,
}

// title : Move the camera, 
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered., 
const orientationV2 cmdDef = 1

type ardrone3CameraOrientationV2 command

func (a ardrone3CameraOrientationV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraOrientationV2 = ardrone3CameraOrientationV2 {
project: ardrone3,
class: camera,
cmd: orientationV2,
}

// title : Move the camera using velocity, 
// desc : Move the camera given velocity consign.\n You can get min and max values for tilt and pan using [CameraVelocityRange](#1-25-4)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered., 
const velocity cmdDef = 2

type ardrone3CameraVelocity command

func (a ardrone3CameraVelocity) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraVelocity = ardrone3CameraVelocity {
project: ardrone3,
class: camera,
cmd: velocity,
}

// Media recording management
const mediaRecord classDef = 7
// title : Take a picture, 
// desc : Take a picture., 
const picture cmdDef = 0

type ardrone3MediaRecordPicture command

func (a ardrone3MediaRecordPicture) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordPicture = ardrone3MediaRecordPicture {
project: ardrone3,
class: mediaRecord,
cmd: picture,
}

// title : Record a video, 
// desc : Record a video., 
const video cmdDef = 1

type ardrone3MediaRecordVideo command

func (a ardrone3MediaRecordVideo) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordVideo = ardrone3MediaRecordVideo {
project: ardrone3,
class: mediaRecord,
cmd: video,
}

// title : Take a picture, 
// desc : Take a picture.\n The type of picture taken is related to the picture setting.\n You can set the picture format by sending the command [SetPictureFormat](#1-19-0). You can also get the current picture format with [PictureFormat](#1-20-0).\n Please note that the time required to take the picture is highly related to this format.\n\n You can check if the picture taking is available with [PictureState](#1-8-2).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after that the picture has been taken)., 
// support : 0901:2.0.1;090c;090e, 
// result : Event [PictureState](#1-8-2) will be triggered with a state busy.\n The drone will take a picture.\n Then, when picture has been taken, notification [PictureEvent](#1-3-0) is triggered.\n And normally [PictureState](#1-8-2) will be triggered with a state ready., 
const pictureV2 cmdDef = 2

type ardrone3MediaRecordPictureV2 command

func (a ardrone3MediaRecordPictureV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordPictureV2 = ardrone3MediaRecordPictureV2 {
project: ardrone3,
class: mediaRecord,
cmd: pictureV2,
}

// title : Record a video, 
// desc : Record a video (or start timelapse).\n You can check if the video recording is available with [VideoState](#1-8-3).\n This command can start a video (obvious huh?), but also a timelapse if the timelapse mode is set. You can check if the timelapse mode is set with the event [TimelapseMode](#1-20-4).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken)., 
// support : 0901:2.0.1;090c;090e, 
// result : The drone will begin or stop to record the video (or timelapse).\n Then, event [VideoState](#1-8-3) will be triggered. Also, notification [VideoEvent](#1-3-1) is triggered., 
const videoV2 cmdDef = 3

type ardrone3MediaRecordVideoV2 command

func (a ardrone3MediaRecordVideoV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordVideoV2 = ardrone3MediaRecordVideoV2 {
project: ardrone3,
class: mediaRecord,
cmd: videoV2,
}

// State of media recording
const mediaRecordState classDef = 8
// title : Picture state, 
// desc : Picture state., 
const pictureStateChanged cmdDef = 0

type ardrone3MediaRecordStatePictureStateChanged command

func (a ardrone3MediaRecordStatePictureStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordStatePictureStateChanged = ardrone3MediaRecordStatePictureStateChanged {
project: ardrone3,
class: mediaRecordState,
cmd: pictureStateChanged,
}

// title : Video record state, 
// desc : Picture record state., 
const videoStateChanged cmdDef = 1

type ardrone3MediaRecordStateVideoStateChanged command

func (a ardrone3MediaRecordStateVideoStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordStateVideoStateChanged = ardrone3MediaRecordStateVideoStateChanged {
project: ardrone3,
class: mediaRecordState,
cmd: videoStateChanged,
}

// title : Picture state, 
// desc : Picture state., 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [TakePicture](#1-7-2) or by a change in the picture state, 
const pictureStateChangedV2 cmdDef = 2

type ardrone3MediaRecordStatePictureStateChangedV2 command

func (a ardrone3MediaRecordStatePictureStateChangedV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordStatePictureStateChangedV2 = ardrone3MediaRecordStatePictureStateChangedV2 {
project: ardrone3,
class: mediaRecordState,
cmd: pictureStateChangedV2,
}

// title : Video record state, 
// desc : Video record state., 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [RecordVideo](#1-7-3) or by a change in the video state, 
const videoStateChangedV2 cmdDef = 3

type ardrone3MediaRecordStateVideoStateChangedV2 command

func (a ardrone3MediaRecordStateVideoStateChangedV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordStateVideoStateChangedV2 = ardrone3MediaRecordStateVideoStateChangedV2 {
project: ardrone3,
class: mediaRecordState,
cmd: videoStateChangedV2,
}

// title : Video resolution, 
// desc : Video resolution.\n Informs about streaming and recording video resolutions.\n Note that this is only an indication about what the resolution should be. To know the real resolution, you should get it from the frame., 
// support : none, 
// triggered : when the resolution changes., 
const videoResolutionState cmdDef = 4

type ardrone3MediaRecordStateVideoResolutionState command

func (a ardrone3MediaRecordStateVideoResolutionState) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordStateVideoResolutionState = ardrone3MediaRecordStateVideoResolutionState {
project: ardrone3,
class: mediaRecordState,
cmd: videoResolutionState,
}

// Events of media recording
const mediaRecordEvent classDef = 3
// title : Picture taken, 
// desc : Picture taken.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**, 
// support : 0901:2.0.1;090c;090e, 
// triggered : after a [TakePicture](#1-7-2), when the picture has been taken (or it has failed)., 
const pictureEventChanged cmdDef = 0

type ardrone3MediaRecordEventPictureEventChanged command

func (a ardrone3MediaRecordEventPictureEventChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordEventPictureEventChanged = ardrone3MediaRecordEventPictureEventChanged {
project: ardrone3,
class: mediaRecordEvent,
cmd: pictureEventChanged,
}

// title : Video record notification, 
// desc : Video record notification.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**, 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [RecordVideo](#1-7-3) or a change in the video state., 
const videoEventChanged cmdDef = 1

type ardrone3MediaRecordEventVideoEventChanged command

func (a ardrone3MediaRecordEventVideoEventChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaRecordEventVideoEventChanged = ardrone3MediaRecordEventVideoEventChanged {
project: ardrone3,
class: mediaRecordEvent,
cmd: videoEventChanged,
}

// State from drone
const pilotingState classDef = 4
// title : Flying state, 
// desc : Flying state., 
// support : 0901;090c;090e, 
// triggered : when the flying state changes., 
const flyingStateChanged cmdDef = 1

type ardrone3PilotingStateFlyingStateChanged command

func (a ardrone3PilotingStateFlyingStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateFlyingStateChanged = ardrone3PilotingStateFlyingStateChanged {
project: ardrone3,
class: pilotingState,
cmd: flyingStateChanged,
}

// title : Alert state, 
// desc : Alert state., 
// support : 0901;090c;090e, 
// triggered : when an alert happens on the drone., 
const alertStateChanged cmdDef = 2

type ardrone3PilotingStateAlertStateChanged command

func (a ardrone3PilotingStateAlertStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateAlertStateChanged = ardrone3PilotingStateAlertStateChanged {
project: ardrone3,
class: pilotingState,
cmd: alertStateChanged,
}

// title : Return home state, 
// desc : Return home state.\n Availability is related to gps fix, magnetometer calibration., 
// support : 0901;090c;090e, 
// triggered : by [ReturnHome](#1-0-5) or when the state of the return home changes., 
const navigateHomeStateChanged cmdDef = 3

type ardrone3PilotingStateNavigateHomeStateChanged command

func (a ardrone3PilotingStateNavigateHomeStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateNavigateHomeStateChanged = ardrone3PilotingStateNavigateHomeStateChanged {
project: ardrone3,
class: pilotingState,
cmd: navigateHomeStateChanged,
}

// title : Drone's position changed, 
// desc : Drone's position changed., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const positionChanged cmdDef = 4

type ardrone3PilotingStatePositionChanged command

func (a ardrone3PilotingStatePositionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStatePositionChanged = ardrone3PilotingStatePositionChanged {
project: ardrone3,
class: pilotingState,
cmd: positionChanged,
}

// title : Drone's speed changed, 
// desc : Drone's speed changed.\n Expressed in the NED referential (North-East-Down)., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const speedChanged cmdDef = 5

type ardrone3PilotingStateSpeedChanged command

func (a ardrone3PilotingStateSpeedChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateSpeedChanged = ardrone3PilotingStateSpeedChanged {
project: ardrone3,
class: pilotingState,
cmd: speedChanged,
}

// title : Drone's attitude changed, 
// desc : Drone's attitude changed., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const attitudeChanged cmdDef = 6

type ardrone3PilotingStateAttitudeChanged command

func (a ardrone3PilotingStateAttitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateAttitudeChanged = ardrone3PilotingStateAttitudeChanged {
project: ardrone3,
class: pilotingState,
cmd: attitudeChanged,
}

// title : Auto takeoff mode, 
// desc : Auto takeoff mode, 
const autoTakeOffModeChanged cmdDef = 7

type ardrone3PilotingStateAutoTakeOffModeChanged command

func (a ardrone3PilotingStateAutoTakeOffModeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateAutoTakeOffModeChanged = ardrone3PilotingStateAutoTakeOffModeChanged {
project: ardrone3,
class: pilotingState,
cmd: autoTakeOffModeChanged,
}

// title : Drone's altitude changed, 
// desc : Drone's altitude changed.\n The altitude reported is the altitude above the take off point.\n To get the altitude above sea level, see [PositionChanged](#1-4-4)., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const altitudeChanged cmdDef = 8

type ardrone3PilotingStateAltitudeChanged command

func (a ardrone3PilotingStateAltitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateAltitudeChanged = ardrone3PilotingStateAltitudeChanged {
project: ardrone3,
class: pilotingState,
cmd: altitudeChanged,
}

// title : Drone's location changed, 
// desc : Drone's location changed.\n This event is meant to replace [PositionChanged](#1-4-4)., 
// support : 0901:4.0.0;090c:4.0.0, 
// triggered : regularly., 
const gpsLocationChanged cmdDef = 9

type ardrone3PilotingStateGpsLocationChanged command

func (a ardrone3PilotingStateGpsLocationChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateGpsLocationChanged = ardrone3PilotingStateGpsLocationChanged {
project: ardrone3,
class: pilotingState,
cmd: gpsLocationChanged,
}

// title : Landing state, 
// desc : Landing state.\n Only available for fixed wings (which have two landing modes)., 
// support : 090e, 
// triggered : when the landing state changes., 
const landingStateChanged cmdDef = 10

type ardrone3PilotingStateLandingStateChanged command

func (a ardrone3PilotingStateLandingStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateLandingStateChanged = ardrone3PilotingStateLandingStateChanged {
project: ardrone3,
class: pilotingState,
cmd: landingStateChanged,
}

// title : Drone's air speed changed, 
// desc : Drone's air speed changed\n Expressed in the drone's referential., 
// support : 090e:1.2.0, 
// triggered : regularly., 
const airSpeedChanged cmdDef = 11

type ardrone3PilotingStateAirSpeedChanged command

func (a ardrone3PilotingStateAirSpeedChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateAirSpeedChanged = ardrone3PilotingStateAirSpeedChanged {
project: ardrone3,
class: pilotingState,
cmd: airSpeedChanged,
}

// title : Move to changed, 
// desc : The drone moves or moved to a given location., 
// support : 090c:4.3.0, 
// triggered : by [MoveTo](#1-0-10) or when the drone did reach the given position., 
const moveToChanged cmdDef = 12

type ardrone3PilotingStatemoveToChanged command

func (a ardrone3PilotingStatemoveToChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStatemoveToChanged = ardrone3PilotingStatemoveToChanged {
project: ardrone3,
class: pilotingState,
cmd: moveToChanged,
}

// title : Motion state, 
// desc : Motion state.\n If [MotionDetection](#1-6-16) is disabled, motion is steady.\n This information is only valid when the drone is not flying., 
// support : 090c:4.3.0, 
// triggered : when the [FlyingState](#1-4-1) is landed and the [MotionDetection](#1-6-16) is enabled and the motion state changes.\n This event is triggered at a filtered rate., 
const motionState cmdDef = 13

type ardrone3PilotingStateMotionState command

func (a ardrone3PilotingStateMotionState) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateMotionState = ardrone3PilotingStateMotionState {
project: ardrone3,
class: pilotingState,
cmd: motionState,
}

// title : Piloted POI state, 
// desc : Piloted POI state., 
// support : 090c:4.3.0, 
// triggered : by [StartPilotedPOI](#1-0-12) or [StopPilotedPOI](#1-0-13) or when piloted POI becomes unavailable., 
const pilotedPOI cmdDef = 14

type ardrone3PilotingStatePilotedPOI command

func (a ardrone3PilotingStatePilotedPOI) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStatePilotedPOI = ardrone3PilotingStatePilotedPOI {
project: ardrone3,
class: pilotingState,
cmd: pilotedPOI,
}

// title : Return home battery capacity, 
// desc : Battery capacity status to return home., 
// support : 090c:4.3.0, 
// triggered : when the status of the battery capacity to do a return home changes. This means that it is triggered either when the battery level changes, when the distance to the home changes or when the position of the home changes., 
const returnHomeBatteryCapacity cmdDef = 15

type ardrone3PilotingStateReturnHomeBatteryCapacity command

func (a ardrone3PilotingStateReturnHomeBatteryCapacity) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateReturnHomeBatteryCapacity = ardrone3PilotingStateReturnHomeBatteryCapacity {
project: ardrone3,
class: pilotingState,
cmd: returnHomeBatteryCapacity,
}

// title : Relative move changed, 
// desc : Relative move changed., 
// triggered : by [MoveRelatively](#1-0-7), or [CancelRelativeMove](#1-0-14) or when the drone's relative move state changes., 
const moveByChanged cmdDef = 16

type ardrone3PilotingStatemoveByChanged command

func (a ardrone3PilotingStatemoveByChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStatemoveByChanged = ardrone3PilotingStatemoveByChanged {
project: ardrone3,
class: pilotingState,
cmd: moveByChanged,
}

// title : Hovering warning, 
// desc : Indicate that the drone may have difficulties to maintain a fix position when hovering., 
// support : 0915, 
// triggered : at connection and on changes., 
const hoveringWarning cmdDef = 17

type ardrone3PilotingStateHoveringWarning command

func (a ardrone3PilotingStateHoveringWarning) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateHoveringWarning = ardrone3PilotingStateHoveringWarning {
project: ardrone3,
class: pilotingState,
cmd: hoveringWarning,
}

// title : Landing auto trigger., 
// desc : Forced landing auto trigger information., 
// support : , 
// triggered : at connection, and when forced landing auto trigger information changes, then every seconds while `reason` is different from `none`., 
const forcedLandingAutoTrigger cmdDef = 18

type ardrone3PilotingStateForcedLandingAutoTrigger command

func (a ardrone3PilotingStateForcedLandingAutoTrigger) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateForcedLandingAutoTrigger = ardrone3PilotingStateForcedLandingAutoTrigger {
project: ardrone3,
class: pilotingState,
cmd: forcedLandingAutoTrigger,
}

// title : Wind state, 
// desc : Wind state., 
// support : 0914, 
// triggered : at connection and on changes., 
const windStateChanged cmdDef = 19

type ardrone3PilotingStateWindStateChanged command

func (a ardrone3PilotingStateWindStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingStateWindStateChanged = ardrone3PilotingStateWindStateChanged {
project: ardrone3,
class: pilotingState,
cmd: windStateChanged,
}

// Events of Piloting
const pilotingEvent classDef = 34
// title : Relative move ended, 
// desc : Relative move ended.\n Informs about the move that the drone managed to do and why it stopped., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : when the drone reaches its target or when it is interrupted by another [moveBy command](#1-0-7) or when an error occurs., 
const moveByEnd cmdDef = 0

type ardrone3PilotingEventmoveByEnd command

func (a ardrone3PilotingEventmoveByEnd) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingEventmoveByEnd = ardrone3PilotingEventmoveByEnd {
project: ardrone3,
class: pilotingEvent,
cmd: moveByEnd,
}

// Network related commands
const network classDef = 13
// title : Scan wifi network, 
// desc : Scan wifi network to get a list of all networks found by the drone, 
// support : 0901;090c;090e, 
// result : Event [WifiScanResults](#1-14-0) is triggered with all networks found.\n When all networks have been sent, event [WifiScanEnded](#1-14-1) is triggered., 
const wifiScan cmdDef = 0

type ardrone3NetworkWifiScan command

func (a ardrone3NetworkWifiScan) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkWifiScan = ardrone3NetworkWifiScan {
project: ardrone3,
class: network,
cmd: wifiScan,
}

// title : Ask for available wifi channels, 
// desc : Ask for available wifi channels.\n The list of available Wifi channels is related to the country of the drone. You can get this country from the event [CountryChanged](#0-3-6)., 
// support : 0901;090c;090e, 
// result : Event [AvailableWifiChannels](#1-14-2) is triggered with all available channels. When all channels have been sent, event [AvailableWifiChannelsCompleted](#1-14-3) is triggered., 
const wifiAuthChannel cmdDef = 1

type ardrone3NetworkWifiAuthChannel command

func (a ardrone3NetworkWifiAuthChannel) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkWifiAuthChannel = ardrone3NetworkWifiAuthChannel {
project: ardrone3,
class: network,
cmd: wifiAuthChannel,
}

// Network state from Product
const networkState classDef = 14
// title : Wifi scan results, 
// desc : Wifi scan results.\n Please note that the list is not complete until you receive the event [WifiScanEnded](#1-14-1)., 
// support : 0901;090c;090e, 
// triggered : for each wifi network scanned after a [ScanWifi](#1-13-0), 
const wifiScanListChanged cmdDef = 0

type ardrone3NetworkStateWifiScanListChanged command

func (a ardrone3NetworkStateWifiScanListChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkStateWifiScanListChanged = ardrone3NetworkStateWifiScanListChanged {
project: ardrone3,
class: networkState,
cmd: wifiScanListChanged,
}

// title : Wifi scan ended, 
// desc : Wifi scan ended.\n When receiving this event, the list of [WifiScanResults](#1-14-0) is complete., 
// support : 0901;090c;090e, 
// triggered : after the last [WifiScanResult](#1-14-0) has been sent., 
const allWifiScanChanged cmdDef = 1

type ardrone3NetworkStateAllWifiScanChanged command

func (a ardrone3NetworkStateAllWifiScanChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkStateAllWifiScanChanged = ardrone3NetworkStateAllWifiScanChanged {
project: ardrone3,
class: networkState,
cmd: allWifiScanChanged,
}

// title : Available wifi channels, 
// desc : Available wifi channels.\n Please note that the list is not complete until you receive the event [AvailableWifiChannelsCompleted](#1-14-3)., 
// support : 0901;090c;090e, 
// triggered : for each available channel after a [GetAvailableWifiChannels](#1-13-1)., 
const wifiAuthChannelListChanged cmdDef = 2

type ardrone3NetworkStateWifiAuthChannelListChanged command

func (a ardrone3NetworkStateWifiAuthChannelListChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkStateWifiAuthChannelListChanged = ardrone3NetworkStateWifiAuthChannelListChanged {
project: ardrone3,
class: networkState,
cmd: wifiAuthChannelListChanged,
}

// title : Available wifi channels completed, 
// desc : Available wifi channels completed.\n When receiving this event, the list of [AvailableWifiChannels](#1-14-2) is complete., 
// support : 0901;090c;090e, 
// triggered : after the last [AvailableWifiChannel](#1-14-2) has been sent., 
const allWifiAuthChannelChanged cmdDef = 3

type ardrone3NetworkStateAllWifiAuthChannelChanged command

func (a ardrone3NetworkStateAllWifiAuthChannelChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkStateAllWifiAuthChannelChanged = ardrone3NetworkStateAllWifiAuthChannelChanged {
project: ardrone3,
class: networkState,
cmd: allWifiAuthChannelChanged,
}

// Piloting Settings commands
const pilotingSettings classDef = 2
// title : Set max altitude, 
// desc : Set max altitude.\n The drone will not fly over this max altitude when it is in manual piloting.\n Please note that if you set a max altitude which is below the current drone altitude, the drone will not go to given max altitude.\n You can get the bounds in the event [MaxAltitude](#1-6-0)., 
// support : 0901;090c;090e, 
// result : The max altitude is set.\n Then, event [MaxAltitude](#1-6-0) is triggered., 
const maxAltitude cmdDef = 0

type ardrone3PilotingSettingsMaxAltitude command

func (a ardrone3PilotingSettingsMaxAltitude) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsMaxAltitude = ardrone3PilotingSettingsMaxAltitude {
project: ardrone3,
class: pilotingSettings,
cmd: maxAltitude,
}

// title : Set max pitch/roll, 
// desc : Set max pitch/roll.\n This represent the max inclination allowed by the drone.\n You can get the bounds with the commands [MaxPitchRoll](#1-6-1)., 
// support : 0901;090c, 
// result : The max pitch/roll is set.\n Then, event [MaxPitchRoll](#1-6-1) is triggered., 
const maxTilt cmdDef = 1

type ardrone3PilotingSettingsMaxTilt command

func (a ardrone3PilotingSettingsMaxTilt) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsMaxTilt = ardrone3PilotingSettingsMaxTilt {
project: ardrone3,
class: pilotingSettings,
cmd: maxTilt,
}

// title : Set absolut control, 
// desc : Set absolut control., 
const absolutControl cmdDef = 2

type ardrone3PilotingSettingsAbsolutControl command

func (a ardrone3PilotingSettingsAbsolutControl) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsAbsolutControl = ardrone3PilotingSettingsAbsolutControl {
project: ardrone3,
class: pilotingSettings,
cmd: absolutControl,
}

// title : Set max distance, 
// desc : Set max distance.\n You can get the bounds from the event [MaxDistance](#1-6-3).\n\n If [Geofence](#1-6-4) is activated, the drone won't fly over the given max distance., 
// support : 0901;090c;090e, 
// result : The max distance is set.\n Then, event [MaxDistance](#1-6-3) is triggered., 
const maxDistance cmdDef = 3

type ardrone3PilotingSettingsMaxDistance command

func (a ardrone3PilotingSettingsMaxDistance) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsMaxDistance = ardrone3PilotingSettingsMaxDistance {
project: ardrone3,
class: pilotingSettings,
cmd: maxDistance,
}

// title : Enable geofence, 
// desc : Enable geofence.\n If geofence is enabled, the drone won't fly over the given max distance.\n You can get the max distance from the event [MaxDistance](#1-6-3). \n For copters: the distance is computed from the controller position, if this position is not known, it will use the take off.\n For fixed wings: the distance is computed from the take off position., 
// support : 0901;090c;090e, 
// result : Geofencing is enabled or disabled.\n Then, event [Geofencing](#1-6-4) is triggered., 
const noFlyOverMaxDistance cmdDef = 4

type ardrone3PilotingSettingsNoFlyOverMaxDistance command

func (a ardrone3PilotingSettingsNoFlyOverMaxDistance) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsNoFlyOverMaxDistance = ardrone3PilotingSettingsNoFlyOverMaxDistance {
project: ardrone3,
class: pilotingSettings,
cmd: noFlyOverMaxDistance,
}

// title : Set autonomous flight max horizontal speed, 
// desc : Set autonomous flight max horizontal speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max horizontal speed is set.\n Then, event [AutonomousFlightMaxHorizontalSpeed](#1-6-5) is triggered., 
const setAutonomousFlightMaxHorizontalSpeed cmdDef = 5

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed command

func (a ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingssetAutonomousFlightMaxHorizontalSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed {
project: ardrone3,
class: pilotingSettings,
cmd: setAutonomousFlightMaxHorizontalSpeed,
}

// title : Set autonomous flight max vertical speed, 
// desc : Set autonomous flight max vertical speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max vertical speed is set.\n Then, event [AutonomousFlightMaxVerticalSpeed](#1-6-6) is triggered., 
const setAutonomousFlightMaxVerticalSpeed cmdDef = 6

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed command

func (a ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingssetAutonomousFlightMaxVerticalSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed {
project: ardrone3,
class: pilotingSettings,
cmd: setAutonomousFlightMaxVerticalSpeed,
}

// title : Set autonomous flight max horizontal acceleration, 
// desc : Set autonomous flight max horizontal acceleration.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max horizontal acceleration is set.\n Then, event [AutonomousFlightMaxHorizontalAcceleration](#1-6-7) is triggered., 
const setAutonomousFlightMaxHorizontalAcceleration cmdDef = 7

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration command

func (a ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration = ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration {
project: ardrone3,
class: pilotingSettings,
cmd: setAutonomousFlightMaxHorizontalAcceleration,
}

// title : Set autonomous flight max vertical acceleration, 
// desc : Set autonomous flight max vertical acceleration.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max vertical acceleration is set.\n Then, event [AutonomousFlightMaxVerticalAcceleration](#1-6-8) is triggered., 
const setAutonomousFlightMaxVerticalAcceleration cmdDef = 8

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration command

func (a ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingssetAutonomousFlightMaxVerticalAcceleration = ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration {
project: ardrone3,
class: pilotingSettings,
cmd: setAutonomousFlightMaxVerticalAcceleration,
}

// title : Set autonomous flight max rotation speed, 
// desc : Set autonomous flight max rotation speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max rotation speed is set.\n Then, event [AutonomousFlightMaxRotationSpeed](#1-6-9) is triggered., 
const setAutonomousFlightMaxRotationSpeed cmdDef = 9

type ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed command

func (a ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingssetAutonomousFlightMaxRotationSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed {
project: ardrone3,
class: pilotingSettings,
cmd: setAutonomousFlightMaxRotationSpeed,
}

// title : Set banked turn mode, 
// desc : Set banked turn mode.\n When banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null., 
// support : 0901:3.2.0;090c:3.2.0, 
// result : The banked turn mode is enabled or disabled.\n Then, event [BankedTurnMode](#1-6-10) is triggered., 
const bankedTurn cmdDef = 10

type ardrone3PilotingSettingsBankedTurn command

func (a ardrone3PilotingSettingsBankedTurn) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsBankedTurn = ardrone3PilotingSettingsBankedTurn {
project: ardrone3,
class: pilotingSettings,
cmd: bankedTurn,
}

// title : Set minimum altitude, 
// desc : Set minimum altitude.\n Only available for fixed wings., 
// support : 090e, 
// result : The minimum altitude is set.\n Then, event [MinimumAltitude](#1-6-11) is triggered., 
const minAltitude cmdDef = 11

type ardrone3PilotingSettingsMinAltitude command

func (a ardrone3PilotingSettingsMinAltitude) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsMinAltitude = ardrone3PilotingSettingsMinAltitude {
project: ardrone3,
class: pilotingSettings,
cmd: minAltitude,
}

// title : Set default circling direction, 
// desc : Set default circling direction. This direction will be used when the drone use an automatic circling or when [CIRCLE](#1-0-9) is sent with direction *default*.\n Only available for fixed wings., 
// support : 090e, 
// result : The circling direction is set.\n Then, event [DefaultCirclingDirection](#1-6-12) is triggered., 
const circlingDirection cmdDef = 12

type ardrone3PilotingSettingsCirclingDirection command

func (a ardrone3PilotingSettingsCirclingDirection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsCirclingDirection = ardrone3PilotingSettingsCirclingDirection {
project: ardrone3,
class: pilotingSettings,
cmd: circlingDirection,
}

// title : Set circling radius, 
// desc : Set circling radius.\n Only available for fixed wings., 
// support : none, 
// result : The circling radius is set.\n Then, event [CirclingRadius](#1-6-13) is triggered., 
const circlingRadius cmdDef = 13

type ardrone3PilotingSettingsCirclingRadius command

func (a ardrone3PilotingSettingsCirclingRadius) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsCirclingRadius = ardrone3PilotingSettingsCirclingRadius {
project: ardrone3,
class: pilotingSettings,
cmd: circlingRadius,
}

// title : Set min circling altitude, 
// desc : Set min circling altitude (not used during take off).\n Only available for fixed wings., 
// support : 090e, 
// result : The circling altitude is set.\n Then, event [CirclingAltitude](#1-6-14) is triggered., 
const circlingAltitude cmdDef = 14

type ardrone3PilotingSettingsCirclingAltitude command

func (a ardrone3PilotingSettingsCirclingAltitude) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsCirclingAltitude = ardrone3PilotingSettingsCirclingAltitude {
project: ardrone3,
class: pilotingSettings,
cmd: circlingAltitude,
}

// title : Set pitch mode, 
// desc : Set pitch mode.\n Only available for fixed wings., 
// support : 090e, 
// result : The pitch mode is set.\n Then, event [PitchMode](#1-6-15) is triggered., 
const pitchMode cmdDef = 15

type ardrone3PilotingSettingsPitchMode command

func (a ardrone3PilotingSettingsPitchMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsPitchMode = ardrone3PilotingSettingsPitchMode {
project: ardrone3,
class: pilotingSettings,
cmd: pitchMode,
}

// title : Enable/disable the motion detection, 
// desc : Enable/disable the motion detection.\n If the motion detection is enabled, the drone will send its [MotionState](#1-4-13) when its [FlyingState](#1-4-1) is landed. If the motion detection is disabled, [MotionState](#1-4-13) is steady., 
// support : 090c:4.3.0, 
// result : The motion detection is enabled or disabled.\n Then, event [MotionDetection](#1-6-16) is triggered. After that, if enabled and [FlyingState](#1-4-1) is landed, the [MotionState](#1-4-13) is triggered upon changes., 
const setMotionDetectionMode cmdDef = 16

type ardrone3PilotingSettingsSetMotionDetectionMode command

func (a ardrone3PilotingSettingsSetMotionDetectionMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsSetMotionDetectionMode = ardrone3PilotingSettingsSetMotionDetectionMode {
project: ardrone3,
class: pilotingSettings,
cmd: setMotionDetectionMode,
}

// Piloting Settings state from product
const pilotingSettingsState classDef = 6
// title : Max altitude, 
// desc : Max altitude.\n The drone will not fly higher than this altitude (above take off point)., 
// support : 0901;090c;090e, 
// triggered : by [SetMaxAltitude](#1-2-0)., 
const maxAltitudeChanged cmdDef = 0

type ardrone3PilotingSettingsStateMaxAltitudeChanged command

func (a ardrone3PilotingSettingsStateMaxAltitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateMaxAltitudeChanged = ardrone3PilotingSettingsStateMaxAltitudeChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: maxAltitudeChanged,
}

// title : Max pitch/roll, 
// desc : Max pitch/roll.\n The drone will not fly higher than this altitude (above take off point)., 
// support : 0901;090c, 
// triggered : by [SetMaxAltitude](#1-2-0)., 
const maxTiltChanged cmdDef = 1

type ardrone3PilotingSettingsStateMaxTiltChanged command

func (a ardrone3PilotingSettingsStateMaxTiltChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateMaxTiltChanged = ardrone3PilotingSettingsStateMaxTiltChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: maxTiltChanged,
}

// title : Absolut control, 
// desc : Absolut control., 
const absolutControlChanged cmdDef = 2

type ardrone3PilotingSettingsStateAbsolutControlChanged command

func (a ardrone3PilotingSettingsStateAbsolutControlChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAbsolutControlChanged = ardrone3PilotingSettingsStateAbsolutControlChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: absolutControlChanged,
}

// title : Max distance, 
// desc : Max distance., 
// support : 0901;090c;090e, 
// triggered : by [SetMaxDistance](#1-2-3)., 
const maxDistanceChanged cmdDef = 3

type ardrone3PilotingSettingsStateMaxDistanceChanged command

func (a ardrone3PilotingSettingsStateMaxDistanceChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateMaxDistanceChanged = ardrone3PilotingSettingsStateMaxDistanceChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: maxDistanceChanged,
}

// title : Geofencing, 
// desc : Geofencing.\n If set, the drone won't fly over the [MaxDistance](#1-6-3)., 
// support : 0901;090c;090e, 
// triggered : by [EnableGeofence](#1-2-4)., 
const noFlyOverMaxDistanceChanged cmdDef = 4

type ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged command

func (a ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateNoFlyOverMaxDistanceChanged = ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: noFlyOverMaxDistanceChanged,
}

// title : Autonomous flight max horizontal speed, 
// desc : Autonomous flight max horizontal speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxHorizontalSpeed](#1-2-5)., 
const autonomousFlightMaxHorizontalSpeed cmdDef = 5

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed command

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed {
project: ardrone3,
class: pilotingSettingsState,
cmd: autonomousFlightMaxHorizontalSpeed,
}

// title : Autonomous flight max vertical speed, 
// desc : Autonomous flight max vertical speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxVerticalSpeed](#1-2-6)., 
const autonomousFlightMaxVerticalSpeed cmdDef = 6

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed command

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAutonomousFlightMaxVerticalSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed {
project: ardrone3,
class: pilotingSettingsState,
cmd: autonomousFlightMaxVerticalSpeed,
}

// title : Autonomous flight max horizontal acceleration, 
// desc : Autonomous flight max horizontal acceleration., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxHorizontalAcceleration](#1-2-7)., 
const autonomousFlightMaxHorizontalAcceleration cmdDef = 7

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration command

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration = ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration {
project: ardrone3,
class: pilotingSettingsState,
cmd: autonomousFlightMaxHorizontalAcceleration,
}

// title : Autonomous flight max vertical acceleration, 
// desc : Autonomous flight max vertical acceleration., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxVerticalAcceleration](#1-2-8)., 
const autonomousFlightMaxVerticalAcceleration cmdDef = 8

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration command

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration = ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration {
project: ardrone3,
class: pilotingSettingsState,
cmd: autonomousFlightMaxVerticalAcceleration,
}

// title : Autonomous flight max rotation speed, 
// desc : Autonomous flight max rotation speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxRotationSpeed](#1-2-9)., 
const autonomousFlightMaxRotationSpeed cmdDef = 9

type ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed command

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateAutonomousFlightMaxRotationSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed {
project: ardrone3,
class: pilotingSettingsState,
cmd: autonomousFlightMaxRotationSpeed,
}

// title : Banked Turn mode, 
// desc : Banked Turn mode.\n If banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null., 
// support : 0901:3.2.0;090c:3.2.0, 
// triggered : by [SetBankedTurnMode](#1-2-10)., 
const bankedTurnChanged cmdDef = 10

type ardrone3PilotingSettingsStateBankedTurnChanged command

func (a ardrone3PilotingSettingsStateBankedTurnChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateBankedTurnChanged = ardrone3PilotingSettingsStateBankedTurnChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: bankedTurnChanged,
}

// title : Min altitude, 
// desc : Min altitude.\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetMinAltitude](#1-2-11)., 
const minAltitudeChanged cmdDef = 11

type ardrone3PilotingSettingsStateMinAltitudeChanged command

func (a ardrone3PilotingSettingsStateMinAltitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateMinAltitudeChanged = ardrone3PilotingSettingsStateMinAltitudeChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: minAltitudeChanged,
}

// title : Circling direction, 
// desc : Circling direction.\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetCirclingDirection](#1-2-12)., 
const circlingDirectionChanged cmdDef = 12

type ardrone3PilotingSettingsStateCirclingDirectionChanged command

func (a ardrone3PilotingSettingsStateCirclingDirectionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateCirclingDirectionChanged = ardrone3PilotingSettingsStateCirclingDirectionChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: circlingDirectionChanged,
}

// title : Circling radius, 
// desc : Circling radius.\n Only sent by fixed wings., 
// support : none, 
// triggered : by [SetCirclingRadius](#1-2-13)., 
const circlingRadiusChanged cmdDef = 13

type ardrone3PilotingSettingsStateCirclingRadiusChanged command

func (a ardrone3PilotingSettingsStateCirclingRadiusChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateCirclingRadiusChanged = ardrone3PilotingSettingsStateCirclingRadiusChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: circlingRadiusChanged,
}

// title : Circling altitude, 
// desc : Circling altitude.\n Bounds will be automatically adjusted according to the [MaxAltitude](#1-6-0).\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetCirclingRadius](#1-2-14) or when bounds change due to [SetMaxAltitude](#1-2-0)., 
const circlingAltitudeChanged cmdDef = 14

type ardrone3PilotingSettingsStateCirclingAltitudeChanged command

func (a ardrone3PilotingSettingsStateCirclingAltitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateCirclingAltitudeChanged = ardrone3PilotingSettingsStateCirclingAltitudeChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: circlingAltitudeChanged,
}

// title : Pitch mode, 
// desc : Pitch mode., 
// support : 090e, 
// triggered : by [SetPitchMode](#1-2-15)., 
const pitchModeChanged cmdDef = 15

type ardrone3PilotingSettingsStatePitchModeChanged command

func (a ardrone3PilotingSettingsStatePitchModeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStatePitchModeChanged = ardrone3PilotingSettingsStatePitchModeChanged {
project: ardrone3,
class: pilotingSettingsState,
cmd: pitchModeChanged,
}

// title : State of the motion detection, 
// desc : State of the motion detection., 
// support : 090c:4.3.0, 
// triggered : by [SetMotionDetectionMode](#1-2-16), 
const motionDetection cmdDef = 16

type ardrone3PilotingSettingsStateMotionDetection command

func (a ardrone3PilotingSettingsStateMotionDetection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pilotingSettingsStateMotionDetection = ardrone3PilotingSettingsStateMotionDetection {
project: ardrone3,
class: pilotingSettingsState,
cmd: motionDetection,
}

// Speed Settings commands
const speedSettings classDef = 11
// title : Set max vertical speed, 
// desc : Set max vertical speed., 
// support : 0901;090c, 
// result : The max vertical speed is set.\n Then, event [MaxVerticalSpeed](#1-12-0) is triggered., 
const maxVerticalSpeed cmdDef = 0

type ardrone3SpeedSettingsMaxVerticalSpeed command

func (a ardrone3SpeedSettingsMaxVerticalSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsMaxVerticalSpeed = ardrone3SpeedSettingsMaxVerticalSpeed {
project: ardrone3,
class: speedSettings,
cmd: maxVerticalSpeed,
}

// title : Set max rotation speed, 
// desc : Set max rotation speed., 
// support : 0901;090c, 
// result : The max rotation speed is set.\n Then, event [MaxRotationSpeed](#1-12-1) is triggered., 
const maxRotationSpeed cmdDef = 1

type ardrone3SpeedSettingsMaxRotationSpeed command

func (a ardrone3SpeedSettingsMaxRotationSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsMaxRotationSpeed = ardrone3SpeedSettingsMaxRotationSpeed {
project: ardrone3,
class: speedSettings,
cmd: maxRotationSpeed,
}

// title : Set the presence of hull protection, 
// desc : Set the presence of hull protection., 
// support : 0901;090c, 
// result : The drone knows that it has a hull protection.\n Then, event [HullProtection](#1-12-2) is triggered., 
const hullProtection cmdDef = 2

type ardrone3SpeedSettingsHullProtection command

func (a ardrone3SpeedSettingsHullProtection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsHullProtection = ardrone3SpeedSettingsHullProtection {
project: ardrone3,
class: speedSettings,
cmd: hullProtection,
}

// title : Set outdoor mode, 
// desc : Set outdoor mode., 
const outdoor cmdDef = 3

type ardrone3SpeedSettingsOutdoor command

func (a ardrone3SpeedSettingsOutdoor) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsOutdoor = ardrone3SpeedSettingsOutdoor {
project: ardrone3,
class: speedSettings,
cmd: outdoor,
}

// title : Set max pitch/roll rotation speed, 
// desc : Set max pitch/roll rotation speed., 
// support : 0901;090c, 
// result : The max pitch/roll rotation speed is set.\n Then, event [MaxPitchRollRotationSpeed](#1-12-4) is triggered., 
const maxPitchRollRotationSpeed cmdDef = 4

type ardrone3SpeedSettingsMaxPitchRollRotationSpeed command

func (a ardrone3SpeedSettingsMaxPitchRollRotationSpeed) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsMaxPitchRollRotationSpeed = ardrone3SpeedSettingsMaxPitchRollRotationSpeed {
project: ardrone3,
class: speedSettings,
cmd: maxPitchRollRotationSpeed,
}

// Speed Settings state from product
const speedSettingsState classDef = 12
// title : Max vertical speed, 
// desc : Max vertical speed., 
// support : 0901;090c, 
// triggered : by [SetMaxVerticalSpeed](#1-11-0)., 
const maxVerticalSpeedChanged cmdDef = 0

type ardrone3SpeedSettingsStateMaxVerticalSpeedChanged command

func (a ardrone3SpeedSettingsStateMaxVerticalSpeedChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsStateMaxVerticalSpeedChanged = ardrone3SpeedSettingsStateMaxVerticalSpeedChanged {
project: ardrone3,
class: speedSettingsState,
cmd: maxVerticalSpeedChanged,
}

// title : Max rotation speed, 
// desc : Max rotation speed., 
// support : 0901;090c, 
// triggered : by [SetMaxRotationSpeed](#1-11-1)., 
const maxRotationSpeedChanged cmdDef = 1

type ardrone3SpeedSettingsStateMaxRotationSpeedChanged command

func (a ardrone3SpeedSettingsStateMaxRotationSpeedChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsStateMaxRotationSpeedChanged = ardrone3SpeedSettingsStateMaxRotationSpeedChanged {
project: ardrone3,
class: speedSettingsState,
cmd: maxRotationSpeedChanged,
}

// title : Presence of hull protection, 
// desc : Presence of hull protection., 
// support : 0901;090c, 
// triggered : by [SetHullProtectionPresence](#1-11-2)., 
const hullProtectionChanged cmdDef = 2

type ardrone3SpeedSettingsStateHullProtectionChanged command

func (a ardrone3SpeedSettingsStateHullProtectionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsStateHullProtectionChanged = ardrone3SpeedSettingsStateHullProtectionChanged {
project: ardrone3,
class: speedSettingsState,
cmd: hullProtectionChanged,
}

// title : Outdoor mode, 
// desc : Outdoor mode., 
const outdoorChanged cmdDef = 3

type ardrone3SpeedSettingsStateOutdoorChanged command

func (a ardrone3SpeedSettingsStateOutdoorChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsStateOutdoorChanged = ardrone3SpeedSettingsStateOutdoorChanged {
project: ardrone3,
class: speedSettingsState,
cmd: outdoorChanged,
}

// title : Max pitch/roll rotation speed, 
// desc : Max pitch/roll rotation speed., 
// support : 0901;090c, 
// triggered : by [SetMaxPitchRollRotationSpeed](#1-11-4)., 
const maxPitchRollRotationSpeedChanged cmdDef = 4

type ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged command

func (a ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var speedSettingsStateMaxPitchRollRotationSpeedChanged = ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged {
project: ardrone3,
class: speedSettingsState,
cmd: maxPitchRollRotationSpeedChanged,
}

// Network settings commands
const networkSettings classDef = 9
// title : Select Wifi, 
// desc : Select or auto-select channel of choosen band., 
// support : 0901;090c;090e, 
// result : The wifi channel changes according to given parameters. Watch out, a disconnection might appear.\n Then, event [WifiSelection](#1-10-0) is triggered., 
const wifiSelection cmdDef = 0

type ardrone3NetworkSettingsWifiSelection command

func (a ardrone3NetworkSettingsWifiSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkSettingsWifiSelection = ardrone3NetworkSettingsWifiSelection {
project: ardrone3,
class: networkSettings,
cmd: wifiSelection,
}

// title : Set wifi security type, 
// desc : Set wifi security type.\n The security will be changed on the next restart, 
// support : 0901;090c;090e, 
// result : The wifi security is set (but not applied until next restart).\n Then, event [WifiSecurityType](#1-10-2) is triggered., 
const wifiSecurity cmdDef = 1

type ardrone3NetworkSettingswifiSecurity command

func (a ardrone3NetworkSettingswifiSecurity) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkSettingswifiSecurity = ardrone3NetworkSettingswifiSecurity {
project: ardrone3,
class: networkSettings,
cmd: wifiSecurity,
}

// Network settings state from product
const networkSettingsState classDef = 10
// title : Wifi selection, 
// desc : Wifi selection., 
// support : 0901;090c;090e, 
// triggered : by [SelectWifi](#1-9-0)., 
const wifiSelectionChanged cmdDef = 0

type ardrone3NetworkSettingsStateWifiSelectionChanged command

func (a ardrone3NetworkSettingsStateWifiSelectionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkSettingsStateWifiSelectionChanged = ardrone3NetworkSettingsStateWifiSelectionChanged {
project: ardrone3,
class: networkSettingsState,
cmd: wifiSelectionChanged,
}

// title : Wifi security type, 
// desc : Wifi security type., 
const wifiSecurityChanged cmdDef = 1

type ardrone3NetworkSettingsStatewifiSecurityChanged command

func (a ardrone3NetworkSettingsStatewifiSecurityChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkSettingsStatewifiSecurityChanged = ardrone3NetworkSettingsStatewifiSecurityChanged {
project: ardrone3,
class: networkSettingsState,
cmd: wifiSecurityChanged,
}

// title : Wifi security type, 
// desc : Wifi security type., 
// support : 0901;090c;090e, 
// triggered : by [SetWifiSecurityType](#1-9-1)., 
const wifiSecurityDUPLICATE cmdDef = 2

type ardrone3NetworkSettingsStatewifiSecurity command

func (a ardrone3NetworkSettingsStatewifiSecurity) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var networkSettingsStatewifiSecurity = ardrone3NetworkSettingsStatewifiSecurity {
project: ardrone3,
class: networkSettingsState,
cmd: wifiSecurity,
}

// Settings state from product
const settingsState classDef = 16
// title : Motor version, 
// desc : Motor version., 
const productMotorVersionListChanged cmdDef = 0

type ardrone3SettingsStateProductMotorVersionListChanged command

func (a ardrone3SettingsStateProductMotorVersionListChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateProductMotorVersionListChanged = ardrone3SettingsStateProductMotorVersionListChanged {
project: ardrone3,
class: settingsState,
cmd: productMotorVersionListChanged,
}

// title : GPS version, 
// desc : GPS version., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const productGPSVersionChanged cmdDef = 1

type ardrone3SettingsStateProductGPSVersionChanged command

func (a ardrone3SettingsStateProductGPSVersionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateProductGPSVersionChanged = ardrone3SettingsStateProductGPSVersionChanged {
project: ardrone3,
class: settingsState,
cmd: productGPSVersionChanged,
}

// title : Motor error, 
// desc : Motor error.\n This event is sent back to *noError* as soon as the motor error disappear. To get the last motor error, see [LastMotorError](#1-16-5), 
// support : 0901;090c;090e, 
// triggered : when a motor error occurs., 
const motorErrorStateChanged cmdDef = 2

type ardrone3SettingsStateMotorErrorStateChanged command

func (a ardrone3SettingsStateMotorErrorStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateMotorErrorStateChanged = ardrone3SettingsStateMotorErrorStateChanged {
project: ardrone3,
class: settingsState,
cmd: motorErrorStateChanged,
}

// title : Motor version, 
// desc : Motor version., 
const motorSoftwareVersionChanged cmdDef = 3

type ardrone3SettingsStateMotorSoftwareVersionChanged command

func (a ardrone3SettingsStateMotorSoftwareVersionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateMotorSoftwareVersionChanged = ardrone3SettingsStateMotorSoftwareVersionChanged {
project: ardrone3,
class: settingsState,
cmd: motorSoftwareVersionChanged,
}

// title : Motor flight status, 
// desc : Motor flight status., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const motorFlightsStatusChanged cmdDef = 4

type ardrone3SettingsStateMotorFlightsStatusChanged command

func (a ardrone3SettingsStateMotorFlightsStatusChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateMotorFlightsStatusChanged = ardrone3SettingsStateMotorFlightsStatusChanged {
project: ardrone3,
class: settingsState,
cmd: motorFlightsStatusChanged,
}

// title : Last motor error, 
// desc : Last motor error.\n This is a reminder of the last error. To know if a motor error is currently happening, see [MotorError](#1-16-2)., 
// support : 0901;090c;090e, 
// triggered : at connection and when an error occurs., 
const motorErrorLastErrorChanged cmdDef = 5

type ardrone3SettingsStateMotorErrorLastErrorChanged command

func (a ardrone3SettingsStateMotorErrorLastErrorChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateMotorErrorLastErrorChanged = ardrone3SettingsStateMotorErrorLastErrorChanged {
project: ardrone3,
class: settingsState,
cmd: motorErrorLastErrorChanged,
}

// title : P7ID, 
// desc : P7ID., 
const p7ID cmdDef = 6

type ardrone3SettingsStateP7ID command

func (a ardrone3SettingsStateP7ID) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateP7ID = ardrone3SettingsStateP7ID {
project: ardrone3,
class: settingsState,
cmd: p7ID,
}

const cPUID cmdDef = 7

type ardrone3SettingsStateCPUID command

func (a ardrone3SettingsStateCPUID) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var settingsStateCPUID = ardrone3SettingsStateCPUID {
project: ardrone3,
class: settingsState,
cmd: cPUID,
}

// Photo settings chosen by the user
const pictureSettings classDef = 19
// title : Set picture format, 
// desc : Set picture format.\n Please note that the time required to take the picture is highly related to this format.\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken)., 
// support : 0901;090c;090e, 
// result : The picture format is set.\n Then, event [PictureFormat](#1-20-0) is triggered., 
const pictureFormatSelection cmdDef = 0

type ardrone3PictureSettingsPictureFormatSelection command

func (a ardrone3PictureSettingsPictureFormatSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsPictureFormatSelection = ardrone3PictureSettingsPictureFormatSelection {
project: ardrone3,
class: pictureSettings,
cmd: pictureFormatSelection,
}

// title : Set White Balance mode, 
// desc : Set White Balance mode., 
// support : 0901;090c;090e, 
// result : The white balance mode is set.\n Then, event [WhiteBalanceMode](#1-20-1) is triggered., 
const autoWhiteBalanceSelection cmdDef = 1

type ardrone3PictureSettingsAutoWhiteBalanceSelection command

func (a ardrone3PictureSettingsAutoWhiteBalanceSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsAutoWhiteBalanceSelection = ardrone3PictureSettingsAutoWhiteBalanceSelection {
project: ardrone3,
class: pictureSettings,
cmd: autoWhiteBalanceSelection,
}

// title : Set image exposure, 
// desc : Set image exposure., 
// support : 0901;090c;090e, 
// result : The exposure is set.\n Then, event [ImageExposure](#1-20-2) is triggered., 
const expositionSelection cmdDef = 2

type ardrone3PictureSettingsExpositionSelection command

func (a ardrone3PictureSettingsExpositionSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsExpositionSelection = ardrone3PictureSettingsExpositionSelection {
project: ardrone3,
class: pictureSettings,
cmd: expositionSelection,
}

// title : Set image saturation, 
// desc : Set image saturation., 
// support : 0901;090c;090e, 
// result : The saturation is set.\n Then, event [ImageSaturation](#1-20-3) is triggered., 
const saturationSelection cmdDef = 3

type ardrone3PictureSettingsSaturationSelection command

func (a ardrone3PictureSettingsSaturationSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsSaturationSelection = ardrone3PictureSettingsSaturationSelection {
project: ardrone3,
class: pictureSettings,
cmd: saturationSelection,
}

// title : Set timelapse mode, 
// desc : Set timelapse mode.\n If timelapse mode is set, instead of taking a video, the drone will take picture regularly.\n Watch out, this command only configure the timelapse mode. Once it is configured, you can start/stop the timelapse with the [RecordVideo](#1-7-3) command., 
// support : 0901;090c;090e, 
// result : The timelapse mode is set (but not started).\n Then, event [TimelapseMode](#1-20-4) is triggered., 
const timelapseSelection cmdDef = 4

type ardrone3PictureSettingsTimelapseSelection command

func (a ardrone3PictureSettingsTimelapseSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsTimelapseSelection = ardrone3PictureSettingsTimelapseSelection {
project: ardrone3,
class: pictureSettings,
cmd: timelapseSelection,
}

// title : Set video autorecord mode, 
// desc : Set video autorecord mode.\n If autorecord is set, video record will be automatically started when the drone takes off and stopped slightly after landing., 
// support : 0901;090c;090e, 
// result : The autorecord mode is set.\n Then, event [AutorecordMode](#1-20-5) is triggered., 
const videoAutorecordSelection cmdDef = 5

type ardrone3PictureSettingsVideoAutorecordSelection command

func (a ardrone3PictureSettingsVideoAutorecordSelection) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsVideoAutorecordSelection = ardrone3PictureSettingsVideoAutorecordSelection {
project: ardrone3,
class: pictureSettings,
cmd: videoAutorecordSelection,
}

// title : Set video stabilization mode, 
// desc : Set video stabilization mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video stabilization mode is set.\n Then, event [VideoStabilizationMode](#1-20-6) is triggered., 
const videoStabilizationMode cmdDef = 6

type ardrone3PictureSettingsVideoStabilizationMode command

func (a ardrone3PictureSettingsVideoStabilizationMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsVideoStabilizationMode = ardrone3PictureSettingsVideoStabilizationMode {
project: ardrone3,
class: pictureSettings,
cmd: videoStabilizationMode,
}

// title : Set video recording mode, 
// desc : Set video recording mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video recording mode is set.\n Then, event [VideoRecordingMode](#1-20-7) is triggered., 
const videoRecordingMode cmdDef = 7

type ardrone3PictureSettingsVideoRecordingMode command

func (a ardrone3PictureSettingsVideoRecordingMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsVideoRecordingMode = ardrone3PictureSettingsVideoRecordingMode {
project: ardrone3,
class: pictureSettings,
cmd: videoRecordingMode,
}

// title : Set video framerate, 
// desc : Set video framerate., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video framerate is set.\n Then, event [VideoFramerate](#1-20-8) is triggered., 
const videoFramerate cmdDef = 8

type ardrone3PictureSettingsVideoFramerate command

func (a ardrone3PictureSettingsVideoFramerate) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsVideoFramerate = ardrone3PictureSettingsVideoFramerate {
project: ardrone3,
class: pictureSettings,
cmd: videoFramerate,
}

// title : Set video resolutions, 
// desc : Set video streaming and recording resolutions., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video resolutions is set.\n Then, event [VideoResolutions](#1-20-9) is triggered., 
const videoResolutions cmdDef = 9

type ardrone3PictureSettingsVideoResolutions command

func (a ardrone3PictureSettingsVideoResolutions) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsVideoResolutions = ardrone3PictureSettingsVideoResolutions {
project: ardrone3,
class: pictureSettings,
cmd: videoResolutions,
}

// Photo settings state from product
const pictureSettingsState classDef = 20
// title : Picture format, 
// desc : Picture format., 
// support : 0901;090c;090e, 
// triggered : by [SetPictureFormat](#1-19-0)., 
const pictureFormatChanged cmdDef = 0

type ardrone3PictureSettingsStatePictureFormatChanged command

func (a ardrone3PictureSettingsStatePictureFormatChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStatePictureFormatChanged = ardrone3PictureSettingsStatePictureFormatChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: pictureFormatChanged,
}

// title : White balance mode, 
// desc : White balance mode., 
// support : 0901;090c;090e, 
// triggered : by [SetWhiteBalanceMode](#1-19-1)., 
const autoWhiteBalanceChanged cmdDef = 1

type ardrone3PictureSettingsStateAutoWhiteBalanceChanged command

func (a ardrone3PictureSettingsStateAutoWhiteBalanceChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateAutoWhiteBalanceChanged = ardrone3PictureSettingsStateAutoWhiteBalanceChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: autoWhiteBalanceChanged,
}

// title : Image exposure, 
// desc : Image exposure., 
// support : 0901;090c;090e, 
// triggered : by [SetImageExposure](#1-19-2)., 
const expositionChanged cmdDef = 2

type ardrone3PictureSettingsStateExpositionChanged command

func (a ardrone3PictureSettingsStateExpositionChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateExpositionChanged = ardrone3PictureSettingsStateExpositionChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: expositionChanged,
}

// title : Image saturation, 
// desc : Image saturation., 
// support : 0901;090c;090e, 
// triggered : by [SetImageSaturation](#1-19-3)., 
const saturationChanged cmdDef = 3

type ardrone3PictureSettingsStateSaturationChanged command

func (a ardrone3PictureSettingsStateSaturationChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateSaturationChanged = ardrone3PictureSettingsStateSaturationChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: saturationChanged,
}

// title : Timelapse mode, 
// desc : Timelapse mode., 
// support : 0901;090c;090e, 
// triggered : by [SetTimelapseMode](#1-19-4)., 
const timelapseChanged cmdDef = 4

type ardrone3PictureSettingsStateTimelapseChanged command

func (a ardrone3PictureSettingsStateTimelapseChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateTimelapseChanged = ardrone3PictureSettingsStateTimelapseChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: timelapseChanged,
}

// title : Video Autorecord mode, 
// desc : Video Autorecord mode., 
// support : 0901;090c;090e, 
// triggered : by [SetVideoAutorecordMode](#1-19-5)., 
const videoAutorecordChanged cmdDef = 5

type ardrone3PictureSettingsStateVideoAutorecordChanged command

func (a ardrone3PictureSettingsStateVideoAutorecordChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateVideoAutorecordChanged = ardrone3PictureSettingsStateVideoAutorecordChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: videoAutorecordChanged,
}

// title : Video stabilization mode, 
// desc : Video stabilization mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoStabilizationMode](#1-19-6)., 
const videoStabilizationModeChanged cmdDef = 6

type ardrone3PictureSettingsStateVideoStabilizationModeChanged command

func (a ardrone3PictureSettingsStateVideoStabilizationModeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateVideoStabilizationModeChanged = ardrone3PictureSettingsStateVideoStabilizationModeChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: videoStabilizationModeChanged,
}

// title : Video recording mode, 
// desc : Video recording mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoRecordingMode](#1-19-7)., 
const videoRecordingModeChanged cmdDef = 7

type ardrone3PictureSettingsStateVideoRecordingModeChanged command

func (a ardrone3PictureSettingsStateVideoRecordingModeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateVideoRecordingModeChanged = ardrone3PictureSettingsStateVideoRecordingModeChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: videoRecordingModeChanged,
}

// title : Video framerate, 
// desc : Video framerate., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoFramerateMode](#1-19-8)., 
const videoFramerateChanged cmdDef = 8

type ardrone3PictureSettingsStateVideoFramerateChanged command

func (a ardrone3PictureSettingsStateVideoFramerateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateVideoFramerateChanged = ardrone3PictureSettingsStateVideoFramerateChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: videoFramerateChanged,
}

// title : Video resolutions, 
// desc : Video resolutions.\n This event informs about the recording AND streaming resolutions., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideResolutions](#1-19-9)., 
const videoResolutionsChanged cmdDef = 9

type ardrone3PictureSettingsStateVideoResolutionsChanged command

func (a ardrone3PictureSettingsStateVideoResolutionsChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pictureSettingsStateVideoResolutionsChanged = ardrone3PictureSettingsStateVideoResolutionsChanged {
project: ardrone3,
class: pictureSettingsState,
cmd: videoResolutionsChanged,
}

// Control media streaming behavior.
const mediaStreaming classDef = 21
// title : Enable/disable video streaming, 
// desc : Enable/disable video streaming., 
// support : 0901;090c;090e, 
// result : The video stream is started or stopped.\n Then, event [VideoStreamState](#1-22-0) is triggered., 
const videoEnable cmdDef = 0

type ardrone3MediaStreamingVideoEnable command

func (a ardrone3MediaStreamingVideoEnable) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaStreamingVideoEnable = ardrone3MediaStreamingVideoEnable {
project: ardrone3,
class: mediaStreaming,
cmd: videoEnable,
}

// title : Set the stream mode, 
// desc : Set the stream mode., 
// support : 0901;090c;090e, 
// result : The stream mode is set.\n Then, event [VideoStreamMode](#1-22-1) is triggered., 
const videoStreamMode cmdDef = 1

type ardrone3MediaStreamingVideoStreamMode command

func (a ardrone3MediaStreamingVideoStreamMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaStreamingVideoStreamMode = ardrone3MediaStreamingVideoStreamMode {
project: ardrone3,
class: mediaStreaming,
cmd: videoStreamMode,
}

// Media streaming status.
const mediaStreamingState classDef = 22
// title : Video stream state, 
// desc : Video stream state., 
// support : 0901;090c;090e, 
// triggered : by [EnableOrDisableVideoStream](#1-21-0)., 
const videoEnableChanged cmdDef = 0

type ardrone3MediaStreamingStateVideoEnableChanged command

func (a ardrone3MediaStreamingStateVideoEnableChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaStreamingStateVideoEnableChanged = ardrone3MediaStreamingStateVideoEnableChanged {
project: ardrone3,
class: mediaStreamingState,
cmd: videoEnableChanged,
}

const videoStreamModeChanged cmdDef = 1

type ardrone3MediaStreamingStateVideoStreamModeChanged command

func (a ardrone3MediaStreamingStateVideoStreamModeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var mediaStreamingStateVideoStreamModeChanged = ardrone3MediaStreamingStateVideoStreamModeChanged {
project: ardrone3,
class: mediaStreamingState,
cmd: videoStreamModeChanged,
}

// GPS settings
const gPSSettings classDef = 23
// title : Set home position, 
// desc : Set home position., 
const setHome cmdDef = 0

type ardrone3GPSSettingsSetHome command

func (a ardrone3GPSSettingsSetHome) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsSetHome = ardrone3GPSSettingsSetHome {
project: ardrone3,
class: gPSSettings,
cmd: setHome,
}

// title : Reset home position, 
// desc : Reset home position., 
// support : 0901;090c, 
// result : The home position is reset.\n Then, event [HomeLocationReset](#1-24-1) is triggered., 
const resetHome cmdDef = 1

type ardrone3GPSSettingsResetHome command

func (a ardrone3GPSSettingsResetHome) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsResetHome = ardrone3GPSSettingsResetHome {
project: ardrone3,
class: gPSSettings,
cmd: resetHome,
}

// title : Set controller gps location, 
// desc : Set controller gps location.\n The user location might be used in case of return home, according to the home type and the accuracy of the given position. You can get the current home type with the event [HomeType](#1-24-4)., 
// support : 0901;090c;090e, 
// result : The controller position is known by the drone.\n Then, event [HomeLocation](#1-24-2) is triggered., 
const sendControllerGPS cmdDef = 2

type ardrone3GPSSettingsSendControllerGPS command

func (a ardrone3GPSSettingsSendControllerGPS) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsSendControllerGPS = ardrone3GPSSettingsSendControllerGPS {
project: ardrone3,
class: gPSSettings,
cmd: sendControllerGPS,
}

// title : Set the preferred home type, 
// desc : Set the preferred home type.\n Please note that this is only a preference. The actual type chosen is given by the event [HomeType](#1-31-2).\n You can get the currently available types with the event [HomeTypeAvailability](#1-31-1)., 
// support : 0901;090c;090e, 
// result : The user choice is known by the drone.\n Then, event [PreferredHomeType](#1-24-4) is triggered., 
const homeType cmdDef = 3

type ardrone3GPSSettingsHomeType command

func (a ardrone3GPSSettingsHomeType) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsHomeType = ardrone3GPSSettingsHomeType {
project: ardrone3,
class: gPSSettings,
cmd: homeType,
}

// title : Set the return home delay, 
// desc : Set the delay after which the drone will automatically try to return home after a disconnection., 
// support : 0901;090c;090e, 
// result : The delay of the return home is set.\n Then, event [ReturnHomeDelay](#1-24-5) is triggered., 
const returnHomeDelay cmdDef = 4

type ardrone3GPSSettingsReturnHomeDelay command

func (a ardrone3GPSSettingsReturnHomeDelay) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsReturnHomeDelay = ardrone3GPSSettingsReturnHomeDelay {
project: ardrone3,
class: gPSSettings,
cmd: returnHomeDelay,
}

// title : Set the return home min altitude, 
// desc : Set the return home minimum altitude. If the drone is below this altitude when starting its return home, it will first reach the minimum altitude. If it is higher than this minimum altitude, it will operate its return home at its actual altitude., 
// support : , 
// result : The minimum altitude for the return home is set.\n Then, event [ReturnHomeMinAltitude](#1-24-7) is triggered., 
const returnHomeMinAltitude cmdDef = 5

type ardrone3GPSSettingsReturnHomeMinAltitude command

func (a ardrone3GPSSettingsReturnHomeMinAltitude) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsReturnHomeMinAltitude = ardrone3GPSSettingsReturnHomeMinAltitude {
project: ardrone3,
class: gPSSettings,
cmd: returnHomeMinAltitude,
}

// GPS settings state
const gPSSettingsState classDef = 24
// title : Home location, 
// desc : Home location., 
// support : 0901;090c;090e, 
// triggered : when [HomeType](#1-31-2) changes. Or by [SetHomeLocation](#1-23-2) when [HomeType](#1-31-2) is Pilot. Or regularly after [SetControllerGPS](#140-1) when [HomeType](#1-31-2) is FollowMeTarget. Or at take off [HomeType](#1-31-2) is Takeoff. Or when the first fix occurs and the [HomeType](#1-31-2) is FirstFix., 
const homeChanged cmdDef = 0

type ardrone3GPSSettingsStateHomeChanged command

func (a ardrone3GPSSettingsStateHomeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateHomeChanged = ardrone3GPSSettingsStateHomeChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: homeChanged,
}

// title : Home location has been reset, 
// desc : Home location has been reset., 
// support : 0901;090c, 
// triggered : by [ResetHomeLocation](#1-23-1)., 
const resetHomeChanged cmdDef = 1

type ardrone3GPSSettingsStateResetHomeChanged command

func (a ardrone3GPSSettingsStateResetHomeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateResetHomeChanged = ardrone3GPSSettingsStateResetHomeChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: resetHomeChanged,
}

// title : Gps fix info, 
// desc : Gps fix info., 
// support : 0901;090c;090e, 
// triggered : on change., 
const gPSFixStateChanged cmdDef = 2

type ardrone3GPSSettingsStateGPSFixStateChanged command

func (a ardrone3GPSSettingsStateGPSFixStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateGPSFixStateChanged = ardrone3GPSSettingsStateGPSFixStateChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: gPSFixStateChanged,
}

// title : Gps update state, 
// desc : Gps update state., 
// support : 0901;090c;090e, 
// triggered : on change., 
const gPSUpdateStateChanged cmdDef = 3

type ardrone3GPSSettingsStateGPSUpdateStateChanged command

func (a ardrone3GPSSettingsStateGPSUpdateStateChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateGPSUpdateStateChanged = ardrone3GPSSettingsStateGPSUpdateStateChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: gPSUpdateStateChanged,
}

// title : Preferred home type, 
// desc : User preference for the home type.\n See [HomeType](#1-31-2) to get the drone actual home type., 
// support : 0901;090c;090e, 
// triggered : by [SetPreferredHomeType](#1-23-3)., 
const homeTypeChanged cmdDef = 4

type ardrone3GPSSettingsStateHomeTypeChanged command

func (a ardrone3GPSSettingsStateHomeTypeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateHomeTypeChanged = ardrone3GPSSettingsStateHomeTypeChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: homeTypeChanged,
}

// title : Return home delay, 
// desc : Return home trigger delay. This delay represents the time after which the return home is automatically triggered after a disconnection., 
// support : 0901;090c;090e, 
// triggered : by [SetReturnHomeDelay](#1-23-4)., 
const returnHomeDelayChanged cmdDef = 5

type ardrone3GPSSettingsStateReturnHomeDelayChanged command

func (a ardrone3GPSSettingsStateReturnHomeDelayChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateReturnHomeDelayChanged = ardrone3GPSSettingsStateReturnHomeDelayChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: returnHomeDelayChanged,
}

// title : Geofence center, 
// desc : Geofence center location. This location represents the center of the geofence zone. This is updated at a maximum frequency of 1 Hz., 
// triggered : when [HomeChanged](#1-24-0) and when [GpsLocationChanged](#1-4-9) before takeoff., 
const geofenceCenterChanged cmdDef = 6

type ardrone3GPSSettingsStateGeofenceCenterChanged command

func (a ardrone3GPSSettingsStateGeofenceCenterChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateGeofenceCenterChanged = ardrone3GPSSettingsStateGeofenceCenterChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: geofenceCenterChanged,
}

// title : Return home min altitude, 
// desc : Minumum altitude for return home changed., 
// triggered : by [SetReturnHomeMinAltitude](#1-23-5)., 
const returnHomeMinAltitudeChanged cmdDef = 7

type ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged command

func (a ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSSettingsStateReturnHomeMinAltitudeChanged = ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged {
project: ardrone3,
class: gPSSettingsState,
cmd: returnHomeMinAltitudeChanged,
}

// Camera state
const cameraState classDef = 25
// title : Camera orientation, 
// desc : Camera orientation., 
// support : 0901;090c;090e, 
// triggered : by [SetCameraOrientation](#1-1-0)., 
const orientationDUPLICATE cmdDef = 0

type ardrone3CameraStateOrientation command

func (a ardrone3CameraStateOrientation) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraStateOrientation = ardrone3CameraStateOrientation {
project: ardrone3,
class: cameraState,
cmd: orientation,
}

// title : Orientation of the camera center, 
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const defaultCameraOrientation cmdDef = 1

type ardrone3CameraStatedefaultCameraOrientation command

func (a ardrone3CameraStatedefaultCameraOrientation) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraStatedefaultCameraOrientation = ardrone3CameraStatedefaultCameraOrientation {
project: ardrone3,
class: cameraState,
cmd: defaultCameraOrientation,
}

// title : Camera orientation, 
// desc : Camera orientation with float arguments., 
// support : 0901;090c;090e, 
// triggered : by [SetCameraOrientationV2](#1-1-1), 
const orientationV2DUPLICATE cmdDef = 2

type ardrone3CameraStateOrientationV2 command

func (a ardrone3CameraStateOrientationV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraStateOrientationV2 = ardrone3CameraStateOrientationV2 {
project: ardrone3,
class: cameraState,
cmd: orientationV2,
}

// title : Orientation of the camera center, 
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const defaultCameraOrientationV2 cmdDef = 3

type ardrone3CameraStatedefaultCameraOrientationV2 command

func (a ardrone3CameraStatedefaultCameraOrientationV2) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraStatedefaultCameraOrientationV2 = ardrone3CameraStatedefaultCameraOrientationV2 {
project: ardrone3,
class: cameraState,
cmd: defaultCameraOrientationV2,
}

// title : Camera velocity range, 
// desc : Camera Orientation velocity limits., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const velocityRange cmdDef = 4

type ardrone3CameraStateVelocityRange command

func (a ardrone3CameraStateVelocityRange) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var cameraStateVelocityRange = ardrone3CameraStateVelocityRange {
project: ardrone3,
class: cameraState,
cmd: velocityRange,
}

// Anti-flickering related commands
const antiflickering classDef = 29
// title : Set the electric frequency, 
// desc : Set the electric frequency of the surrounding lights.\n This is used to avoid the video flickering in auto mode. You can get the current antiflickering mode with the event [AntiflickeringModeChanged](#1-30-1)., 
// support : 0901;090c, 
// result : The electric frequency is set.\n Then, event [ElectricFrequency](#1-30-0) is triggered., 
const electricFrequency cmdDef = 0

type ardrone3AntiflickeringelectricFrequency command

func (a ardrone3AntiflickeringelectricFrequency) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var antiflickeringelectricFrequency = ardrone3AntiflickeringelectricFrequency {
project: ardrone3,
class: antiflickering,
cmd: electricFrequency,
}

// title : Set the antiflickering mode, 
// desc : Set the antiflickering mode.\n If auto, the drone will detect when flickers appears on the video and trigger the antiflickering.\n In this case, this electric frequency it will use will be the one specified in the event [ElectricFrequency](#1-29-0).\n Forcing the antiflickering (FixedFiftyHertz or FixedFiftyHertz) can reduce luminosity of the video., 
// support : 0901;090c, 
// result : The antiflickering mode is set.\n Then, event [AntiflickeringMode](#1-30-1) is triggered., 
const setMode cmdDef = 1

type ardrone3AntiflickeringsetMode command

func (a ardrone3AntiflickeringsetMode) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var antiflickeringsetMode = ardrone3AntiflickeringsetMode {
project: ardrone3,
class: antiflickering,
cmd: setMode,
}

// Anti-flickering related states
const antiflickeringState classDef = 30
// title : Electric frequency, 
// desc : Electric frequency.\n This piece of information is used for the antiflickering when the [AntiflickeringMode](#1-30-1) is set to *auto*., 
// support : 0901;090c, 
// triggered : by [SetElectricFrequency](#1-29-0)., 
const electricFrequencyChanged cmdDef = 0

type ardrone3AntiflickeringStateelectricFrequencyChanged command

func (a ardrone3AntiflickeringStateelectricFrequencyChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var antiflickeringStateelectricFrequencyChanged = ardrone3AntiflickeringStateelectricFrequencyChanged {
project: ardrone3,
class: antiflickeringState,
cmd: electricFrequencyChanged,
}

// title : Antiflickering mode, 
// desc : Antiflickering mode., 
// support : 0901;090c, 
// triggered : by [SetAntiflickeringMode](#1-29-1)., 
const modeChanged cmdDef = 1

type ardrone3AntiflickeringStatemodeChanged command

func (a ardrone3AntiflickeringStatemodeChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var antiflickeringStatemodeChanged = ardrone3AntiflickeringStatemodeChanged {
project: ardrone3,
class: antiflickeringState,
cmd: modeChanged,
}

// GPS related States
const gPSState classDef = 31
// title : Number of GPS satellites, 
// desc : Number of GPS satellites., 
// support : 0901;090c;090e, 
// triggered : on change., 
const numberOfSatelliteChanged cmdDef = 0

type ardrone3GPSStateNumberOfSatelliteChanged command

func (a ardrone3GPSStateNumberOfSatelliteChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSStateNumberOfSatelliteChanged = ardrone3GPSStateNumberOfSatelliteChanged {
project: ardrone3,
class: gPSState,
cmd: numberOfSatelliteChanged,
}

// title : Home type availability, 
// desc : Home type availability., 
// support : 0901;090c;090e, 
// triggered : when the availability of, at least, one type changes.\n This might be due to controller position availability, gps fix before take off or other reason., 
const homeTypeAvailabilityChanged cmdDef = 1

type ardrone3GPSStateHomeTypeAvailabilityChanged command

func (a ardrone3GPSStateHomeTypeAvailabilityChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSStateHomeTypeAvailabilityChanged = ardrone3GPSStateHomeTypeAvailabilityChanged {
project: ardrone3,
class: gPSState,
cmd: homeTypeAvailabilityChanged,
}

// title : Home type, 
// desc : Home type.\n This choice is made by the drone, according to the [PreferredHomeType](#1-24-4) and the [HomeTypeAvailability](#1-31-1). The drone will choose the type matching with the user preference only if this type is available. If not, it will chose a type in this order:\n FOLLOWEE ; TAKEOFF ; PILOT ; FIRST_FIX, 
// support : 0901;090c;090e, 
// triggered : when the return home type chosen by the drone changes.\n This might be produced by a user preference triggered by [SetPreferedHomeType](#1-23-3) or by a change in the [HomeTypesAvailabilityChanged](#1-31-1)., 
const homeTypeChosenChanged cmdDef = 2

type ardrone3GPSStateHomeTypeChosenChanged command

func (a ardrone3GPSStateHomeTypeChosenChanged) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var gPSStateHomeTypeChosenChanged = ardrone3GPSStateHomeTypeChosenChanged {
project: ardrone3,
class: gPSState,
cmd: homeTypeChosenChanged,
}

// Pro features enabled on the Bebop
const pROState classDef = 32
// title : Pro features, 
// desc : Pro features., 
const features cmdDef = 0

type ardrone3PROStateFeatures command

func (a ardrone3PROStateFeatures) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var pROStateFeatures = ardrone3PROStateFeatures {
project: ardrone3,
class: pROState,
cmd: features,
}

// Information about the connected accessories
const accessoryState classDef = 33
// title : List of connected accessories, 
// desc : List of all connected accessories. This event presents the list of all connected accessories. To actually use the component, use the component dedicated feature., 
// support : 090e:1.5.0, 
// triggered : at connection or when an accessory is connected., 
const connectedAccessories cmdDef = 0

type ardrone3AccessoryStateConnectedAccessories command

func (a ardrone3AccessoryStateConnectedAccessories) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var accessoryStateConnectedAccessories = ardrone3AccessoryStateConnectedAccessories {
project: ardrone3,
class: accessoryState,
cmd: connectedAccessories,
}

// title : Connected accessories battery, 
// desc : Connected accessories battery., 
// support : none, 
const battery cmdDef = 1

type ardrone3AccessoryStateBattery command

func (a ardrone3AccessoryStateBattery) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var accessoryStateBattery = ardrone3AccessoryStateBattery {
project: ardrone3,
class: accessoryState,
cmd: battery,
}

// Sounds related commands
const sound classDef = 35
// title : Start alert sound, 
// desc : Start the alert sound. The alert sound can only be started when the drone is not flying., 
// support : none, 
// result : The drone makes a sound and send back [AlertSoundState](#1-36-0) with state playing., 
const startAlertSound cmdDef = 0

type ardrone3SoundStartAlertSound command

func (a ardrone3SoundStartAlertSound) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var soundStartAlertSound = ardrone3SoundStartAlertSound {
project: ardrone3,
class: sound,
cmd: startAlertSound,
}

// title : Stop alert sound, 
// desc : Stop the alert sound., 
// support : none, 
// result : The drone stops its alert sound and send back [AlertSoundState](#1-36-0) with state stopped., 
const stopAlertSound cmdDef = 1

type ardrone3SoundStopAlertSound command

func (a ardrone3SoundStopAlertSound) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var soundStopAlertSound = ardrone3SoundStopAlertSound {
project: ardrone3,
class: sound,
cmd: stopAlertSound,
}

// Sounds related events
const soundState classDef = 36
// title : Alert sound state, 
// desc : Alert sound state., 
// support : none, 
// triggered : by [StartAlertSound](#1-35-0) or [StopAlertSound](#1-35-1) or when the drone starts or stops to play an alert sound by itself., 
const alertSound cmdDef = 0

type ardrone3SoundStateAlertSound command

func (a ardrone3SoundStateAlertSound) decode() {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
}

var soundStateAlertSound = ardrone3SoundStateAlertSound {
project: ardrone3,
class: soundState,
cmd: alertSound,
}

type decoder interface {
decode()
}

var commandMap = map[command]decoder {
command(pilotingTakeOff) : pilotingTakeOff,
command(pilotingPCMD) : pilotingPCMD,
command(pilotingLanding) : pilotingLanding,
command(pilotingEmergency) : pilotingEmergency,
command(pilotingNavigateHome) : pilotingNavigateHome,
command(pilotingAutoTakeOffMode) : pilotingAutoTakeOffMode,
command(pilotingmoveBy) : pilotingmoveBy,
command(pilotingUserTakeOff) : pilotingUserTakeOff,
command(pilotingCircle) : pilotingCircle,
command(pilotingmoveTo) : pilotingmoveTo,
command(pilotingCancelMoveTo) : pilotingCancelMoveTo,
command(pilotingStartPilotedPOI) : pilotingStartPilotedPOI,
command(pilotingStopPilotedPOI) : pilotingStopPilotedPOI,
command(pilotingCancelMoveBy) : pilotingCancelMoveBy,
command(animationsFlip) : animationsFlip,
command(cameraOrientation) : cameraOrientation,
command(cameraOrientationV2) : cameraOrientationV2,
command(cameraVelocity) : cameraVelocity,
command(mediaRecordPicture) : mediaRecordPicture,
command(mediaRecordVideo) : mediaRecordVideo,
command(mediaRecordPictureV2) : mediaRecordPictureV2,
command(mediaRecordVideoV2) : mediaRecordVideoV2,
command(mediaRecordStatePictureStateChanged) : mediaRecordStatePictureStateChanged,
command(mediaRecordStateVideoStateChanged) : mediaRecordStateVideoStateChanged,
command(mediaRecordStatePictureStateChangedV2) : mediaRecordStatePictureStateChangedV2,
command(mediaRecordStateVideoStateChangedV2) : mediaRecordStateVideoStateChangedV2,
command(mediaRecordStateVideoResolutionState) : mediaRecordStateVideoResolutionState,
command(mediaRecordEventPictureEventChanged) : mediaRecordEventPictureEventChanged,
command(mediaRecordEventVideoEventChanged) : mediaRecordEventVideoEventChanged,
command(pilotingStateFlyingStateChanged) : pilotingStateFlyingStateChanged,
command(pilotingStateAlertStateChanged) : pilotingStateAlertStateChanged,
command(pilotingStateNavigateHomeStateChanged) : pilotingStateNavigateHomeStateChanged,
command(pilotingStatePositionChanged) : pilotingStatePositionChanged,
command(pilotingStateSpeedChanged) : pilotingStateSpeedChanged,
command(pilotingStateAttitudeChanged) : pilotingStateAttitudeChanged,
command(pilotingStateAutoTakeOffModeChanged) : pilotingStateAutoTakeOffModeChanged,
command(pilotingStateAltitudeChanged) : pilotingStateAltitudeChanged,
command(pilotingStateGpsLocationChanged) : pilotingStateGpsLocationChanged,
command(pilotingStateLandingStateChanged) : pilotingStateLandingStateChanged,
command(pilotingStateAirSpeedChanged) : pilotingStateAirSpeedChanged,
command(pilotingStatemoveToChanged) : pilotingStatemoveToChanged,
command(pilotingStateMotionState) : pilotingStateMotionState,
command(pilotingStatePilotedPOI) : pilotingStatePilotedPOI,
command(pilotingStateReturnHomeBatteryCapacity) : pilotingStateReturnHomeBatteryCapacity,
command(pilotingStatemoveByChanged) : pilotingStatemoveByChanged,
command(pilotingStateHoveringWarning) : pilotingStateHoveringWarning,
command(pilotingStateForcedLandingAutoTrigger) : pilotingStateForcedLandingAutoTrigger,
command(pilotingStateWindStateChanged) : pilotingStateWindStateChanged,
command(pilotingEventmoveByEnd) : pilotingEventmoveByEnd,
command(networkWifiScan) : networkWifiScan,
command(networkWifiAuthChannel) : networkWifiAuthChannel,
command(networkStateWifiScanListChanged) : networkStateWifiScanListChanged,
command(networkStateAllWifiScanChanged) : networkStateAllWifiScanChanged,
command(networkStateWifiAuthChannelListChanged) : networkStateWifiAuthChannelListChanged,
command(networkStateAllWifiAuthChannelChanged) : networkStateAllWifiAuthChannelChanged,
command(pilotingSettingsMaxAltitude) : pilotingSettingsMaxAltitude,
command(pilotingSettingsMaxTilt) : pilotingSettingsMaxTilt,
command(pilotingSettingsAbsolutControl) : pilotingSettingsAbsolutControl,
command(pilotingSettingsMaxDistance) : pilotingSettingsMaxDistance,
command(pilotingSettingsNoFlyOverMaxDistance) : pilotingSettingsNoFlyOverMaxDistance,
command(pilotingSettingssetAutonomousFlightMaxHorizontalSpeed) : pilotingSettingssetAutonomousFlightMaxHorizontalSpeed,
command(pilotingSettingssetAutonomousFlightMaxVerticalSpeed) : pilotingSettingssetAutonomousFlightMaxVerticalSpeed,
command(pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) : pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration,
command(pilotingSettingssetAutonomousFlightMaxVerticalAcceleration) : pilotingSettingssetAutonomousFlightMaxVerticalAcceleration,
command(pilotingSettingssetAutonomousFlightMaxRotationSpeed) : pilotingSettingssetAutonomousFlightMaxRotationSpeed,
command(pilotingSettingsBankedTurn) : pilotingSettingsBankedTurn,
command(pilotingSettingsMinAltitude) : pilotingSettingsMinAltitude,
command(pilotingSettingsCirclingDirection) : pilotingSettingsCirclingDirection,
command(pilotingSettingsCirclingRadius) : pilotingSettingsCirclingRadius,
command(pilotingSettingsCirclingAltitude) : pilotingSettingsCirclingAltitude,
command(pilotingSettingsPitchMode) : pilotingSettingsPitchMode,
command(pilotingSettingsSetMotionDetectionMode) : pilotingSettingsSetMotionDetectionMode,
command(pilotingSettingsStateMaxAltitudeChanged) : pilotingSettingsStateMaxAltitudeChanged,
command(pilotingSettingsStateMaxTiltChanged) : pilotingSettingsStateMaxTiltChanged,
command(pilotingSettingsStateAbsolutControlChanged) : pilotingSettingsStateAbsolutControlChanged,
command(pilotingSettingsStateMaxDistanceChanged) : pilotingSettingsStateMaxDistanceChanged,
command(pilotingSettingsStateNoFlyOverMaxDistanceChanged) : pilotingSettingsStateNoFlyOverMaxDistanceChanged,
command(pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) : pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed,
command(pilotingSettingsStateAutonomousFlightMaxVerticalSpeed) : pilotingSettingsStateAutonomousFlightMaxVerticalSpeed,
command(pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) : pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration,
command(pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) : pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration,
command(pilotingSettingsStateAutonomousFlightMaxRotationSpeed) : pilotingSettingsStateAutonomousFlightMaxRotationSpeed,
command(pilotingSettingsStateBankedTurnChanged) : pilotingSettingsStateBankedTurnChanged,
command(pilotingSettingsStateMinAltitudeChanged) : pilotingSettingsStateMinAltitudeChanged,
command(pilotingSettingsStateCirclingDirectionChanged) : pilotingSettingsStateCirclingDirectionChanged,
command(pilotingSettingsStateCirclingRadiusChanged) : pilotingSettingsStateCirclingRadiusChanged,
command(pilotingSettingsStateCirclingAltitudeChanged) : pilotingSettingsStateCirclingAltitudeChanged,
command(pilotingSettingsStatePitchModeChanged) : pilotingSettingsStatePitchModeChanged,
command(pilotingSettingsStateMotionDetection) : pilotingSettingsStateMotionDetection,
command(speedSettingsMaxVerticalSpeed) : speedSettingsMaxVerticalSpeed,
command(speedSettingsMaxRotationSpeed) : speedSettingsMaxRotationSpeed,
command(speedSettingsHullProtection) : speedSettingsHullProtection,
command(speedSettingsOutdoor) : speedSettingsOutdoor,
command(speedSettingsMaxPitchRollRotationSpeed) : speedSettingsMaxPitchRollRotationSpeed,
command(speedSettingsStateMaxVerticalSpeedChanged) : speedSettingsStateMaxVerticalSpeedChanged,
command(speedSettingsStateMaxRotationSpeedChanged) : speedSettingsStateMaxRotationSpeedChanged,
command(speedSettingsStateHullProtectionChanged) : speedSettingsStateHullProtectionChanged,
command(speedSettingsStateOutdoorChanged) : speedSettingsStateOutdoorChanged,
command(speedSettingsStateMaxPitchRollRotationSpeedChanged) : speedSettingsStateMaxPitchRollRotationSpeedChanged,
command(networkSettingsWifiSelection) : networkSettingsWifiSelection,
command(networkSettingswifiSecurity) : networkSettingswifiSecurity,
command(networkSettingsStateWifiSelectionChanged) : networkSettingsStateWifiSelectionChanged,
command(networkSettingsStatewifiSecurityChanged) : networkSettingsStatewifiSecurityChanged,
command(networkSettingsStatewifiSecurity) : networkSettingsStatewifiSecurity,
command(settingsStateProductMotorVersionListChanged) : settingsStateProductMotorVersionListChanged,
command(settingsStateProductGPSVersionChanged) : settingsStateProductGPSVersionChanged,
command(settingsStateMotorErrorStateChanged) : settingsStateMotorErrorStateChanged,
command(settingsStateMotorSoftwareVersionChanged) : settingsStateMotorSoftwareVersionChanged,
command(settingsStateMotorFlightsStatusChanged) : settingsStateMotorFlightsStatusChanged,
command(settingsStateMotorErrorLastErrorChanged) : settingsStateMotorErrorLastErrorChanged,
command(settingsStateP7ID) : settingsStateP7ID,
command(settingsStateCPUID) : settingsStateCPUID,
command(pictureSettingsPictureFormatSelection) : pictureSettingsPictureFormatSelection,
command(pictureSettingsAutoWhiteBalanceSelection) : pictureSettingsAutoWhiteBalanceSelection,
command(pictureSettingsExpositionSelection) : pictureSettingsExpositionSelection,
command(pictureSettingsSaturationSelection) : pictureSettingsSaturationSelection,
command(pictureSettingsTimelapseSelection) : pictureSettingsTimelapseSelection,
command(pictureSettingsVideoAutorecordSelection) : pictureSettingsVideoAutorecordSelection,
command(pictureSettingsVideoStabilizationMode) : pictureSettingsVideoStabilizationMode,
command(pictureSettingsVideoRecordingMode) : pictureSettingsVideoRecordingMode,
command(pictureSettingsVideoFramerate) : pictureSettingsVideoFramerate,
command(pictureSettingsVideoResolutions) : pictureSettingsVideoResolutions,
command(pictureSettingsStatePictureFormatChanged) : pictureSettingsStatePictureFormatChanged,
command(pictureSettingsStateAutoWhiteBalanceChanged) : pictureSettingsStateAutoWhiteBalanceChanged,
command(pictureSettingsStateExpositionChanged) : pictureSettingsStateExpositionChanged,
command(pictureSettingsStateSaturationChanged) : pictureSettingsStateSaturationChanged,
command(pictureSettingsStateTimelapseChanged) : pictureSettingsStateTimelapseChanged,
command(pictureSettingsStateVideoAutorecordChanged) : pictureSettingsStateVideoAutorecordChanged,
command(pictureSettingsStateVideoStabilizationModeChanged) : pictureSettingsStateVideoStabilizationModeChanged,
command(pictureSettingsStateVideoRecordingModeChanged) : pictureSettingsStateVideoRecordingModeChanged,
command(pictureSettingsStateVideoFramerateChanged) : pictureSettingsStateVideoFramerateChanged,
command(pictureSettingsStateVideoResolutionsChanged) : pictureSettingsStateVideoResolutionsChanged,
command(mediaStreamingVideoEnable) : mediaStreamingVideoEnable,
command(mediaStreamingVideoStreamMode) : mediaStreamingVideoStreamMode,
command(mediaStreamingStateVideoEnableChanged) : mediaStreamingStateVideoEnableChanged,
command(mediaStreamingStateVideoStreamModeChanged) : mediaStreamingStateVideoStreamModeChanged,
command(gPSSettingsSetHome) : gPSSettingsSetHome,
command(gPSSettingsResetHome) : gPSSettingsResetHome,
command(gPSSettingsSendControllerGPS) : gPSSettingsSendControllerGPS,
command(gPSSettingsHomeType) : gPSSettingsHomeType,
command(gPSSettingsReturnHomeDelay) : gPSSettingsReturnHomeDelay,
command(gPSSettingsReturnHomeMinAltitude) : gPSSettingsReturnHomeMinAltitude,
command(gPSSettingsStateHomeChanged) : gPSSettingsStateHomeChanged,
command(gPSSettingsStateResetHomeChanged) : gPSSettingsStateResetHomeChanged,
command(gPSSettingsStateGPSFixStateChanged) : gPSSettingsStateGPSFixStateChanged,
command(gPSSettingsStateGPSUpdateStateChanged) : gPSSettingsStateGPSUpdateStateChanged,
command(gPSSettingsStateHomeTypeChanged) : gPSSettingsStateHomeTypeChanged,
command(gPSSettingsStateReturnHomeDelayChanged) : gPSSettingsStateReturnHomeDelayChanged,
command(gPSSettingsStateGeofenceCenterChanged) : gPSSettingsStateGeofenceCenterChanged,
command(gPSSettingsStateReturnHomeMinAltitudeChanged) : gPSSettingsStateReturnHomeMinAltitudeChanged,
command(cameraStateOrientation) : cameraStateOrientation,
command(cameraStatedefaultCameraOrientation) : cameraStatedefaultCameraOrientation,
command(cameraStateOrientationV2) : cameraStateOrientationV2,
command(cameraStatedefaultCameraOrientationV2) : cameraStatedefaultCameraOrientationV2,
command(cameraStateVelocityRange) : cameraStateVelocityRange,
command(antiflickeringelectricFrequency) : antiflickeringelectricFrequency,
command(antiflickeringsetMode) : antiflickeringsetMode,
command(antiflickeringStateelectricFrequencyChanged) : antiflickeringStateelectricFrequencyChanged,
command(antiflickeringStatemodeChanged) : antiflickeringStatemodeChanged,
command(gPSStateNumberOfSatelliteChanged) : gPSStateNumberOfSatelliteChanged,
command(gPSStateHomeTypeAvailabilityChanged) : gPSStateHomeTypeAvailabilityChanged,
command(gPSStateHomeTypeChosenChanged) : gPSStateHomeTypeChosenChanged,
command(pROStateFeatures) : pROStateFeatures,
command(accessoryStateConnectedAccessories) : accessoryStateConnectedAccessories,
command(accessoryStateBattery) : accessoryStateBattery,
command(soundStartAlertSound) : soundStartAlertSound,
command(soundStopAlertSound) : soundStopAlertSound,
command(soundStateAlertSound) : soundStateAlertSound,
}

