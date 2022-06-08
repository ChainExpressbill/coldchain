import React from 'react';
import ModalPortal from './ModalPortal';
import styles from './modal.module.scss';

interface ConfirmModalProps {
  msg: string;
  confirmAction: () => void;
  closeModal: () => void;
}

const ConfirmModal = ({
  msg,
  confirmAction,
  closeModal,
}: ConfirmModalProps) => {
  return (
    <ModalPortal>
      <div className={styles.Modal__Background}>
        <div className={styles.Modal__Container}>
          <div className="text-xl text-black">{msg}</div>
          <div className="flex gap-10">
            <button className={styles.Modal__Button} onClick={confirmAction}>
              확인
            </button>
            <button className={styles.Modal__Button} onClick={closeModal}>
              취소
            </button>
          </div>
        </div>
      </div>
    </ModalPortal>
  );
};

export default ConfirmModal;
