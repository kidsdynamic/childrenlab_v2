export class Dashboard {
  signup: DashboardSignup[];
  activity: DashboardActivity[];
}

class DashboardSignup {
  signupCount: number;
  date: string;
}
class DashboardActivity {
  activityCount: number;
  date: string;
}
