import React from 'react';
import { Header, Footer } from 'components/layout';
import { APP_STAGE } from 'constants/environment';
import {
  LoginPage,
  JoinPage,
  DashBoardPage,
  OrderPage,
  OrderListPage,
} from 'pages';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

if (APP_STAGE === 'local') {
  require('./mocks');
}

function App() {
  return (
    <div className="flex flex-col w-full h-screen">
      <Header />
      <main className="mt-24 h-full">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<LoginPage />} />
            <Route path="/join" element={<JoinPage />} />
            <Route path="/dashboard" element={<DashBoardPage />} />
            <Route path="/order/:id?" element={<OrderPage />} />
            <Route path="/orders" element={<OrderListPage />} />
          </Routes>
        </BrowserRouter>
      </main>
      <Footer />
    </div>
  );
}

export default App;
