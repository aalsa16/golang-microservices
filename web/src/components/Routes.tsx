import { Navigate, Outlet, RouterProvider, createBrowserRouter } from "react-router-dom";
import { useAuth } from "../context/AuthProvider";
import Login from "./Login";
import Signup from "./Signup";
import GetStarted from "./GetStarted";
import Home from "./Home";
import History from "./History";

const ProtectedRoute = () => {
    // @ts-ignore
    const { accessToken } = useAuth();
  
    // Check if the user is authenticated
    if (!accessToken) {
      // If not authenticated, redirect to the login page
      return <Navigate to="/getstarted" />;
    }
  
    // If authenticated, render the child routes
    return <Outlet />;
};

const Routes = () => {
    // @ts-ignore
  const { accessToken } = useAuth();

  console.log(accessToken);

  // Define routes accessible only to authenticated users
  const routesForAuthenticatedOnly = [
    {
      path: "/",
      element: <ProtectedRoute />, // Wrap the component in ProtectedRoute
      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/history",
          element: <History />,
        },
      ],
    },
  ];

  // Define routes accessible only to non-authenticated users
  const routesForNotAuthenticatedOnly = [
    {
        path: "/getstarted",
        element: <GetStarted />
    },
    {
      path: "/signup",
      element: <Signup />,
    },
    {
      path: "/login",
      element: <Login />,
    },
  ];

  // Combine and conditionally include routes based on authentication status
  const router = createBrowserRouter([
    ...(!accessToken ? routesForNotAuthenticatedOnly : []),
    ...routesForAuthenticatedOnly,
  ]);

  // Provide the router configuration using RouterProvider
  return <RouterProvider router={router} />;
};

export default Routes;