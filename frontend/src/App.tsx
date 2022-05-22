import { APP_STAGE } from 'constants/environment';

if (APP_STAGE === 'local') {
  require('./mocks');
}

function App() {
  return <div>Cold Chain</div>;
}

export default App;
