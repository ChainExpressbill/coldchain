import React from 'react';
import { Link, LinkProps, useMatch, useResolvedPath } from 'react-router-dom';
import classnames from 'classnames';
import styles from './menu.module.scss';
import { MdDashboard, MdOutlineCreate, MdList } from 'react-icons/md';

function CustomLink({ children, to, ...props }: LinkProps) {
  const resolved = useResolvedPath(to);
  const match = useMatch({ path: resolved.pathname, end: true });
  const activeClassName = classnames(match ? styles.active : '');

  return (
    <div>
      <Link className={activeClassName} to={to} {...props}>
        {children}
      </Link>
    </div>
  );
}

function SideMenu() {
  const asideClassnames = classnames('w-full flex flex-col', styles.lnb);

  return (
    <aside className={asideClassnames}>
      <div className="Logo__Box px-6 py-4 cursor-pointer border-r-[1px] border-main">
        <Link to="/dashboard">
          <img src="/logo_long_white.jpg" alt="로고 이미지" />
        </Link>
      </div>
      <ul className={styles.Menu__List}>
        <li className={styles.Menu__Title}>
          <CustomLink to="/dashboard">
            <MdDashboard className="mr-2" />
            현황 보드
          </CustomLink>
        </li>
        <li className={styles.Menu__Title}>
          <CustomLink to="/orders/new">
            <MdOutlineCreate className="mr-2" />
            주문 등록
          </CustomLink>
        </li>
        <li className={styles.Menu__Title}>
          <CustomLink to="/orders">
            <MdList className="mr-2" />
            주문 조회
          </CustomLink>
        </li>
      </ul>
    </aside>
  );
}

export default SideMenu;
