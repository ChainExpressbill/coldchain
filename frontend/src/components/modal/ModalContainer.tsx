import { ReactNode } from 'react';
import ModalPortal from './ModalPortal';

interface ModalConstainerProps {
  children: ReactNode;
}

const ModalContainer = ({ children }: ModalConstainerProps) => {
  return (
    <ModalPortal>
      <div>
        <div>{children}</div>
      </div>
    </ModalPortal>
  );
};

export default ModalContainer;
