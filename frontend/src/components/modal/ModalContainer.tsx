import React, { ReactNode } from 'react';
import ModalPortal from './ModalPortal';

interface ModalConstainerProps {
  children: ReactNode;
}

const ModalContainer = ({ children }: ModalConstainerProps) => {
  return (
    <ModalPortal>
      <div className="h-full w-full fixed flex items-center justify-center left-0 top-0 text-center bg-black bg-opacity-25">
        <div className="h-1/3 w-1/3 bg-white rounded-md flex items-center justify-center flex-col gap-6">
          {children}
        </div>
      </div>
    </ModalPortal>
  );
};

export default ModalContainer;
