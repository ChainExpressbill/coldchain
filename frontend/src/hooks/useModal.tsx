import { useState } from 'react';

const useModal = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [modalMsg, setModalMsg] = useState('');
  const [modalType, setModalType] = useState('');
  return { isOpen, setIsOpen, modalMsg, setModalMsg, modalType, setModalType };
};

export default useModal;
