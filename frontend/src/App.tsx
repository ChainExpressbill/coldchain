import React from 'react';
import { Header, Footer } from 'components/layout';
import { APP_STAGE } from 'constants/environment';
import { SignIn, SignUp, DashBoard, ManageOrder, ViewOrder } from 'pages';
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
            <Route path="/" element={<SignIn />} />
            <Route path="/signup" element={<SignUp />} />
            <Route path="/dashboard" element={<DashBoard />} />
            <Route path="/manageorder" element={<ManageOrder />} />
            <Route path="/vieworder" element={<ViewOrder />} />
          </Routes>
        </BrowserRouter>
      </main>
      <Footer />
    </div>
  );
}

export default App;
