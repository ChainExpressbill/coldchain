import { rest } from 'msw';
import { API_URL } from 'constants/environment';

const commonHandlers = [
  rest.get(`${API_URL}`, (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ key: 'value' }));
  }),
];

export default commonHandlers;
