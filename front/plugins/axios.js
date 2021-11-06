export default function({ $axios }) {
    $axios.interceptors.request.use(request => {
    // console.log("Starting Request: ", request);
        return request;
    });
}
