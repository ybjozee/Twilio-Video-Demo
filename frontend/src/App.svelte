<script>
    import 'izitoast/dist/css/iziToast.min.css'
    import 'izitoast/dist/js/iziToast.min'

    import {connectToRoom} from "./Helper";

    let hasJoinedRoom = false;

    const handleSubmit = async e => {
        const formData = new FormData(e.target);
        const roomName = formData.get('roomName');
        hasJoinedRoom = true
        const videoContainer = document.getElementById('remote-media');
        await connectToRoom(roomName, videoContainer)
    }
</script>

<main>
    <h1> Svelte Go Twilio Video Chat App </h1>
    <div id="remote-media"></div>
    {#if !hasJoinedRoom}
        <form on:submit|preventDefault={handleSubmit}>
            <div>
                <label for="roomName">Room name</label>
                <input
                        type="text"
                        id="roomName"
                        name="roomName"
                        value=""
                />
            </div>
            <button type="submit">Submit</button>
        </form>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }

    @media (min-width: 640px) {
        main {
            max-width: none;
        }
    }
</style>