import React from 'react';
import { APP_STAGE } from 'constants/environment';
import {
  LoginPage,
  JoinPage,
  DashBoardPage,
  OrderPage,
  OrderListPage,
} from 'pages';
import {
  BrowserRouter,
  Route,
  Routes,
  Navigate,
  useLocation,
} from 'react-router-dom';

if (APP_STAGE === 'local') {
  require('./mocks');
}

function RequireAuth({ children }: { children: JSX.Element }) {
  // const auth = useAuth();
  const auth = { user: 'jtpark' };
  const location = useLocation();

  if (!auth.user) {
    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to="/" state={{ from: location }} replace />;
  }

  return children;
}

function App() {
  return (
    <div className="flex flex-col w-full h-screen">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LoginPage />} />
          <Route path="/join" element={<JoinPage />} />
          <Route
            path="dashboard"
            element={
              <RequireAuth>
                <DashBoardPage />
              </RequireAuth>
            }
          />
          <Route path="order/:id?" element={<OrderPage />} />
          <Route path="orders" element={<OrderListPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
