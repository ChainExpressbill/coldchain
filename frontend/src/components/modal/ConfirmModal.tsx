import React from 'react';
import ModalContainer from './ModalContainer';

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
    <ModalContainer>
      <div className="text-2xl text-black">{msg}</div>
      <div className="flex gap-10">
        <button
          className="w-16 h-12 border border-black rounded-md text-black"
          onClick={confirmAction}
        >
          확인
        </button>
        <button
          className="w-16 h-12 border border-black rounded-md text-black"
          onClick={closeModal}
        >
          취소
        </button>
      </div>
    </ModalContainer>
  );
};

export default ConfirmModal;
