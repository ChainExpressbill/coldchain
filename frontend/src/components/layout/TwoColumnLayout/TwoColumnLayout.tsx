import React from 'react';
import { Outlet } from 'react-router-dom';
import SideMenu from 'components/side';
import Header from '../Header';
import 'react-loading-skeleton/dist/skeleton.css';

function TwoColumnLayout() {
  return (
    <section className="h-full grid grid__Two__Column__Layout">
      <SideMenu />
      <article>
        <Header />
        <div className="p-6 h-[calc(100%_-_3rem)]">
          <Outlet />
        </div>
      </article>
    </section>
  );
}

export default TwoColumnLayout;
