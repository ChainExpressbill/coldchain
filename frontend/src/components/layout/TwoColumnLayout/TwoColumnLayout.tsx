import React from 'react';
import { Outlet } from 'react-router-dom';
import classnames from 'classnames';
import styles from './layout.module.scss';
import SideMenu from 'components/side';
import Header from '../Header';

function TwoColumnLayout() {
  const cls = classnames('h-full grid', styles.grid__Two__Column);
  return (
    <section className={cls}>
      <SideMenu />
      <article>
        <Header />
        <div className="p-6">
          <Outlet />
        </div>
      </article>
    </section>
  );
}

export default TwoColumnLayout;
