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
import { TwoColumnLayout } from 'components/layout';

if (APP_STAGE === 'local') {
  require('./mocks');
}

function RequireAuth({ children }: { children: JSX.Element }) {
  // const auth = useAuth();
  const auth = { user: 'jtpark' };
  const location = useLocation();

  if (!auth.user) {
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
          <Route element={<TwoColumnLayout />}>
            <Route
              path="/dashboard"
              element={
                <RequireAuth>
                  <DashBoardPage />
                </RequireAuth>
              }
            />
            <Route path="/order/:id?" element={<OrderPage />} />
            <Route path="/orders" element={<OrderListPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
