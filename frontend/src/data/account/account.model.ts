export interface LoginForm {
  id: string;
  password: string;
}

export interface JoinForm extends LoginForm {
  name: string;
  emailAddress: string;
}
