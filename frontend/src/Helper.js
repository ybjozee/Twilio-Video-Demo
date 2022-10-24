import {connect, createLocalVideoTrack} from "twilio-video";
import axios from "axios";
import iziToast from "izitoast";

const notify = message => {
    iziToast.success({
        message,
        position: 'topRight'

    });
}

const axiosInstance = axios.create({
    baseURL: "http://localhost:8000",
});

const getAccessToken = async (roomName) => {
    const response = await axiosInstance.post("", {roomName});
    const {jwt} = response.data;
    return jwt;
};

export const connectToRoom = async (roomName, videoContainer) => {
    const token = await getAccessToken(roomName);
    const room = await connect(token, {name: roomName});
    const videoTrack = await createLocalVideoTrack();
    notify(`Successfully joined a Room: ${room}`);
    videoContainer.appendChild(videoTrack.attach());
    room.on('participantConnected', participant => {
        notify(`A remote participant connected: ${participant.identity}`);
        participant.tracks.forEach(publication => {
            if (publication.isSubscribed) {
                const track = publication.track;
                videoContainer.appendChild(track.attach());
            }
        });
        participant.on('trackSubscribed', track => {
            videoContainer.appendChild(track.attach());
        });
    });
    room.participants.forEach(participant => {
        participant.tracks.forEach(publication => {
            if (publication.track) {
                videoContainer.appendChild(publication.track.attach());
            }
        });
        participant.on('trackSubscribed', track => {
            videoContainer.appendChild(track.attach());
        });
    });
}