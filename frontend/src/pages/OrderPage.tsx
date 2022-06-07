import React, { useEffect } from 'react';
import useModal from 'hooks/useModal';
import { useForm } from 'react-hook-form';
import { AlertModal, ConfirmModal, ModalPortal } from 'components/modal';
import { OrderRequestBody } from 'data/order/order.model';
import {
  useOrderCreate,
  useOrderDetail,
  useOrderUpdate,
} from 'data/order/order.hooks';
import { AxiosError } from 'axios';
import { FormInput } from 'components/form';
import { useParams } from 'react-router-dom';

function OrderPage() {
  const { id } = useParams();
  const orderId = parseInt(id ?? '-1'); //  TODO: 처리를 다르게 해야하나?

  const {
    register,
    handleSubmit,
    formState: { errors: formErrors, isValid: isFormValid },
    getValues,
    reset,
  } = useForm<OrderRequestBody>({ mode: 'onChange' });
  const { refetch } = useOrderDetail(orderId, {
    enabled: false,
  });

  useEffect(() => {
    if (orderId !== -1) {
      refetch().then((response) => {
        if (response.data !== undefined) {
          const targetOrderData = {
            orderer: response.data.orderer,
            receiver: response.data.receiver,
            drugName: response.data.drugName,
            quantity: response.data.quantity,
            drugStandard: response.data.drugStandard,
            storageCondition: response.data.storageCondition,
            oid: response.data.oid,
            id: response.data.id,
          };
          reset(targetOrderData);
        }
      });
    }
  }, [orderId]);

  const {
    isOpen: confirmModalIsOpen,
    setIsOpen: setConfirmModalIsOpen,
    modalMsg: confirmModalMsg,
    setModalMsg: setConfirmModalMsg,
  } = useModal();
  const {
    isOpen: alertModalIsOpen,
    setIsOpen: setAlertModalIsOpen,
    modalMsg: alertModalMsg,
    setModalMsg: setAlertModalMsg,
  } = useModal();

  const { mutate: useCreateOrder } = useOrderCreate({
    onSuccess(data) {
      console.log('success', data);
    },
    onError(orderError: AxiosError) {
      if (orderError.response !== undefined) {
        console.log(orderError.response);
        setAlertModalMsg(() => '주문 내용이 잘못 입력되었습니다.');
      }
      setAlertModalIsOpen(() => true);
    },
  });

  const { mutate: useUpdateOrder } = useOrderUpdate({
    onSuccess(data) {
      console.log('success', data);
    },
    onError(orderError: AxiosError) {
      if (orderError.response !== undefined) {
        console.log(orderError.response);
        setAlertModalMsg(() => '주문 내용이 잘못 입력되었습니다.');
      }
      setAlertModalIsOpen(() => true);
    },
  });

  const confirmCreateOrder = () => {
    //  TODO: accountId, registerName 처리 필요
    const data: OrderRequestBody = {
      ...getValues(),
      accountId: 'jtpark',
      registerName: 'testman',
    };
    useCreateOrder(data);
    setConfirmModalIsOpen(() => false);
  };

  const confirmUpdateOrder = () => {
    const data: OrderRequestBody = {
      ...getValues(),
      accountId: 'jtpark',
      registerName: 'testman',
    };
    useUpdateOrder(data);
    setConfirmModalIsOpen(() => false);
  };

  const onCreateOrderSubmit = () => {
    setConfirmModalMsg('주문을 등록하시겠습니까?');
    setConfirmModalIsOpen(() => true);
  };

  const onUpdateOrderSubmit = () => {
    setConfirmModalMsg('주문 내용을 수정하시겠습니까?');
    setConfirmModalIsOpen(() => true);
  };

  return (
    <div className="text-black h-full flex flex-col align-items-center justify-center gap-12">
      <form
        onSubmit={
          orderId === -1
            ? handleSubmit(onCreateOrderSubmit)
            : handleSubmit(onUpdateOrderSubmit)
        }
        className="w-3/4 ml-auto mr-auto grid grid-cols-1 gap-6 Rounded__Shadow__Border p-6"
      >
        <FormInput
          register={register}
          valueLabel="주문자"
          valueName={'orderer'}
          formError={formErrors.orderer}
          placeholder="주문하시는 분의 이름을 입력해주세요"
          validMsg="주문하시는 분의 이름을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="수령인"
          valueName={'receiver'}
          formError={formErrors.receiver}
          placeholder="수령하시는 분의 이름을 입력해주세요"
          validMsg="수령하시는 분의 이름을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="약품명"
          valueName={'drugName'}
          formError={formErrors.drugName}
          placeholder="약품명을 입력해주세요"
          validMsg="약품명을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="약품 수량"
          valueName={'quantity'}
          formError={formErrors.quantity}
          placeholder="약품 수량을 입력해주세요"
          validMsg="약품 수량을 확인해주세요"
          type="number"
        />
        <FormInput
          register={register}
          valueLabel="약품 규격"
          valueName={'drugStandard'}
          formError={formErrors.drugStandard}
          placeholder="약품 규격을 입력해주세요"
          validMsg="약품 규격을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="약품 보관 조건"
          valueName={'storageCondition'}
          formError={formErrors.storageCondition}
          placeholder="약품 보관 조건을 입력해주세요"
          validMsg="약품 보관 조건을 확인해주세요"
        />
        <button
          className="mr-auto ml-auto w-1/4 p-2 border border-gray-300 disabled:text-white disabled:bg-red-500 disabled:cursor-not-allowed"
          type="submit"
          disabled={!isFormValid}
        >
          주문등록
        </button>
      </form>
      <ModalPortal>
        {confirmModalIsOpen && (
          <ConfirmModal
            msg={confirmModalMsg}
            confirmAction={
              orderId === -1 ? confirmCreateOrder : confirmUpdateOrder
            }
            closeModal={() => {
              setConfirmModalIsOpen(() => false);
            }}
          />
        )}
        {alertModalIsOpen && (
          <AlertModal
            msg={alertModalMsg}
            type="error"
            closeModal={() => {
              setAlertModalIsOpen(() => false);
            }}
          />
        )}
      </ModalPortal>
    </div>
  );
}

export default OrderPage;
