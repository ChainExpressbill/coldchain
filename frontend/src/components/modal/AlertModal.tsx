import React from 'react';
import ModalPortal from './ModalPortal';
import styles from './modal.module.scss';

interface AlertModalProps {
  msg: string;
  type: string;
  closeModal: () => void;
}

const AlertModal = ({ msg, type, closeModal }: AlertModalProps) => {
  const msgStyle = `text-xl ${
    type === 'success' ? 'text-black' : 'text-red-500'
  }`;
  return (
    <ModalPortal>
      <div className={styles.Modal__Background} onClick={closeModal}>
        <div
          className={styles.Modal__Container}
          onClick={(e) => e.stopPropagation()}
        >
          <div className={msgStyle}>{msg}</div>
          <button className={styles.Modal__Button} onClick={closeModal}>
            확인
          </button>
        </div>
      </div>
    </ModalPortal>
  );
};

export default AlertModal;
