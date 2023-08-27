import axios from "axios";

const customFetch = axios.create({
    baseURL: "http://127.0.0.1:8080/",
    headers: {
        "Content-type": "application/json",
    },
    withCredentials: true,
});

const refreshToken = async () => {
    try {
        const refreshToken: any = localStorage.getItem("refresh_token");

        const response = await customFetch.post('auth/refreshToken', {
            refresh_token: refreshToken,
        }, {
            headers: {
                'Content-Type': 'application/json'
            }
        });

        // check if refresh token is expired, if it is then delete and redo auth

        return response.data.message.access_token;
    } catch (e) {
        console.log("Error", e);
    }
};

customFetch.interceptors.response.use(
    (response) => {
        return response;
    },
    async function (error) {
        const originalRequest = error.config;
        if (error.response.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;

            const resp = await refreshToken();

            localStorage.setItem('access_token', resp);
            customFetch.defaults.headers.common[
                "Authorization"
            ] = `Bearer ${resp}`;
            return customFetch(originalRequest);
        }
        return Promise.reject(error);
    }
);

export default customFetch;