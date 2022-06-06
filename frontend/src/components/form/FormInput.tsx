import React from 'react';
import { FieldError, Path, UseFormRegister } from 'react-hook-form';

interface FormInputProps<T> {
  register: UseFormRegister<T>;
  valueLabel: string;
  valueName: Path<T>;
  formError: FieldError | undefined;
  placeholder: string;
  validMsg: string;
  type?: string;
}

const FormInput = <T,>({
  register,
  valueLabel,
  valueName,
  formError,
  placeholder,
  validMsg,
  type = 'text',
}: FormInputProps<T>) => {
  return (
    <div className="Order__Form__Container__Box">
      <div className="Order__Form__Label__Box">
        <label>{valueLabel}</label>
      </div>
      <input
        className="Order__Form__Input__Box"
        placeholder={placeholder}
        {...register(valueName, {
          required: {
            value: true,
            message: validMsg,
          },
          valueAsNumber: type === 'number',
        })}
        type={type}
      />
      {formError && (
        <div className="Order__Form__ValidMsg__Box">
          <span className="text-red-600 text-sm">{formError.message}</span>
        </div>
      )}
    </div>
  );
};

export default FormInput;
