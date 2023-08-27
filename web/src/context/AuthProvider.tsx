import { createContext, useContext, useEffect, useMemo, useState } from "react";
import customFetch from "../api/axios";

const AuthContext = createContext({});

const AuthProvider = ({ children }: any) => {
    // State to hold the authentication token
    const [accessToken, setAccessToken] = useState(localStorage.getItem("access_token"));

    // Function to set the authentication token
    const setToken = (accessToken: string) => {
        setAccessToken(accessToken);
    };

    useEffect(() => {
        if (accessToken) {
            customFetch.defaults.headers.common["Authorization"] = "Bearer " + accessToken;
            localStorage.setItem('access_token', accessToken);
        } else {
            delete customFetch.defaults.headers.common["Authorization"];
            localStorage.removeItem('access_token')
        }
    }, [accessToken]);

    // Memoized value of the authentication context
    const contextValue = useMemo(
        () => ({
            accessToken,
            setToken,
        }),
        [accessToken]
    );

    // Provide the authentication context to the children components
    return (
        <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
    );
};

export const useAuth = () => {
    return useContext(AuthContext);
};

export default AuthProvider;