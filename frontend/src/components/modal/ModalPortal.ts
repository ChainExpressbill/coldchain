import { ReactNode } from 'react';
import reactDom from 'react-dom';

interface ModalPortalProps {
  children: ReactNode;
}

const ModalPortal = ({ children }: ModalPortalProps) => {
  const modalRoot = document.getElementById('modal-root') as HTMLElement;
  return reactDom.createPortal(children, modalRoot);
};

export default ModalPortal;
