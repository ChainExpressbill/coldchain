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
          <Route
            element={
              <RequireAuth>
                <TwoColumnLayout />
              </RequireAuth>
            }
          >
            <Route path="/dashboard" element={<DashBoardPage />} />
            <Route path="/orders">
              <Route index element={<OrderListPage />} />
              <Route path="new" element={<OrderPage />} />
              <Route path=":id" element={<OrderPage />} />
            </Route>
          </Route>
          <Route path="*" element={<div>Not Found</div>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
