import React, { useEffect } from 'react';
import useModal from 'hooks/useModal';
import { useForm } from 'react-hook-form';
import { AlertModal, ConfirmModal } from 'components/modal';
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
  const orderId = Number(id);
  const isOrderDetailView = !isNaN(orderId);

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
    if (isOrderDetailView) {
      refetch().then((response) => {
        if (response.data) {
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
    } else {
      reset({
        orderer: '',
        receiver: '',
        drugName: '',
        quantity: 0,
        drugStandard: '',
        storageCondition: '',
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
    setConfirmModalIsOpen(false);
  };

  const confirmUpdateOrder = () => {
    const data: OrderRequestBody = {
      ...getValues(),
      accountId: 'jtpark',
      registerName: 'testman',
    };
    useUpdateOrder(data);
    setConfirmModalIsOpen(false);
  };

  const onCreateOrderSubmit = () => {
    setConfirmModalMsg('주문을 등록하시겠습니까?');
    setConfirmModalIsOpen(true);
  };

  const onUpdateOrderSubmit = () => {
    setConfirmModalMsg('주문 내용을 수정하시겠습니까?');
    setConfirmModalIsOpen(true);
  };

  return (
    <div className="text-black h-full flex flex-col align-items-center justify-center">
      <form
        onSubmit={
          !isOrderDetailView
            ? handleSubmit(onCreateOrderSubmit)
            : handleSubmit(onUpdateOrderSubmit)
        }
        className="w-1/2 mx-auto grid gap-6 Rounded__Shadow__Border p-6"
      >
        {isOrderDetailView && <div>주문 상세 내역</div>}
        <FormInput
          register={register}
          valueLabel="주문 업체"
          valueName={'orderer'}
          formError={formErrors.orderer}
          placeholder="주문 업체의 이름을 입력해주세요"
          validMsg="주문 업체의 분의 이름을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="수령 업체"
          valueName={'receiver'}
          formError={formErrors.receiver}
          placeholder="수령 업체의 이름을 입력해주세요"
          validMsg="수령 업체의 이름을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="제품명"
          valueName={'drugName'}
          formError={formErrors.drugName}
          placeholder="제품명을 입력해주세요"
          validMsg="제품명을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="제품 수량"
          valueName={'quantity'}
          formError={formErrors.quantity}
          placeholder="제품 수량을 입력해주세요"
          validMsg="제품 수량을 확인해주세요"
          type="number"
        />
        <FormInput
          register={register}
          valueLabel="제품 규격"
          valueName={'drugStandard'}
          formError={formErrors.drugStandard}
          placeholder="제품 규격을 입력해주세요"
          validMsg="제품 규격을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="보관 조건"
          valueName={'storageCondition'}
          formError={formErrors.storageCondition}
          placeholder="보관 조건을 입력해주세요"
          validMsg="보관 조건을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="배송 기사 이름"
          valueName={'deliveryDriverName'}
          formError={formErrors.deliveryDriverName}
          placeholder="배송 기사 이름을 입력해주세요"
          validMsg="배송 기사 이름을 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="배송 기사 연락처"
          valueName={'deliveryDriverTelNo'}
          formError={formErrors.deliveryDriverTelNo}
          placeholder="배송 기사 연락처를 입력해주세요"
          validMsg="배송 기사 연락처를 확인해주세요"
        />
        <FormInput
          register={register}
          valueLabel="메모"
          valueName={'memo'}
          formError={formErrors.memo}
          placeholder="메모를 입력해주세요"
          validMsg="메모를 확인해주세요"
        />
        <button
          className="mx-auto w-1/4 p-2 border border-gray-300 text-white bg-red-500 disabled:opacity-60 disabled:cursor-not-allowed"
          type="submit"
          disabled={!isFormValid}
        >
          {isOrderDetailView ? '주문 수정' : '주문 등록'}
        </button>
      </form>
      {confirmModalIsOpen && (
        <ConfirmModal
          msg={confirmModalMsg}
          confirmAction={
            isOrderDetailView ? confirmUpdateOrder : confirmCreateOrder
          }
          closeModal={() => {
            setConfirmModalIsOpen(false);
          }}
        />
      )}
      {alertModalIsOpen && (
        <AlertModal
          msg={alertModalMsg}
          type="error"
          closeModal={() => {
            setAlertModalIsOpen(false);
          }}
        />
      )}
    </div>
  );
}

export default OrderPage;
