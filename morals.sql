create database morals;

\c morals;

create table morals (
  id serial,
  category char(15),
  quote text,
  created_on timestamp without time zone
);

create table metrics ( 
  id serial,
  moral_id int,
  created_on timestamp without time zone
);

insert into morals ( category, quote, created_on ) values
  ( 'philosophy','Great leaders inspire greatness in others.', now()),
  ( 'philosophy','Belief is not a matter of choice, but of conviction.',now()),
  ( 'philosophy','Easy is the path to wisdom for those not blinded by ego',now()),
  ( 'philosophy','Easy is the path to wisdom for those not blinded by themselves',now()),
  ( 'philosophy','A plan is only as good as those who see it through.',now()),
  ( 'philosophy','The best confidence builder is experience.',now()),
  ( 'philosophy','Trust in your friends, and they''ll have reason to trust in you.',now()),
  ( 'philosophy','You hold onto friends by keeping your heart a little softer than your head.',now()),
  ( 'philosophy','Heroes are made by the times.',now()),
  ( 'philosophy','Ignore your instincts at your peril.',now()),
  ( 'philosophy','Most powerful is he who controls his own power.',now()),
  ( 'philosophy','The winding path to peace is always a worthy one, regardless of how many turns it takes.',now()),
  ( 'philosophy','Fail with honor rather than succeed by fraud.',now()),
  ( 'philosophy','Greed and fear of loss are the roots that lead to the tree of evil.',now()),
  ( 'philosophy','When surrounded by war, one must eventually choose a side.',now()),
  ( 'philosophy','Arrogance diminishes wisdom.',now()),
  ( 'philosophy','Truth enlightens the mind, but won''t always bring happiness to your heart.',now()),
  ( 'philosophy','Fear is a disease; hope is its only cure.',now()),
  ( 'philosophy','A single chance is a galaxy of hope.',now()),
  ( 'philosophy','It is a rough road that leads to the heights of greatness.',now()),
  ( 'philosophy','The costs of war can never be truly accounted for.',now()),
  ( 'philosophy','Compromise is a virtue to be cultivated, not a weakness to be despised.',now()),
  ( 'philosophy','A secret shared is a trust formed.',now()),
  ( 'philosophy','Easy is the path to wisdom for those not blinded by themselves',now()),
  ( 'philosophy','A lesson learned is a lesson earned.',now()),
  ( 'philosophy','Overconfidence is the most dangerous form of carelessness.',now()),
  ( 'philosophy','The first step to correcting a mistake is patience.',now()),
  ( 'philosophy','A true heart should never be doubted.',now()),
  ( 'philosophy','Believe in yourself or no one else will.',now()),
  ( 'philosophy','No gift is more precious than trust.',now()),
  ( 'philosophy','Sometimes, accepting help is harder than offering it.',now()),
  ( 'philosophy','Attachment is not compassion.',now()),
  ( 'philosophy','For everything you gain, you lose something else.',now()),
  ( 'philosophy','It is the quest for honor that makes one honorable.',now()),
  ( 'philosophy','Easy isn''t always simple.',now()),
  ( 'philosophy','If you ignore the past, you jeopardize your future.',now()),
  ( 'philosophy','Fear not for the future, weep not for the past.',now()),
  ( 'philosophy','In war, truth is the first casualty.',now()),
  ( 'philosophy','Searching for the truth is easy. Accepting the truth is hard.',now()),
  ( 'philosophy','A wise leader knows when to follow.',now()),
  ( 'philosophy','Courage makes heroes, but trust builds friendship.',now()),
  ( 'philosophy','Choose what is right, not what is easy',now()),
  ( 'philosophy','The most dangerous beast is the beast within.',now()),
  ( 'philosophy','Who my father was matters less than my memory of him.',now()),
  ( 'philosophy','Adversity is friendship''s truest test.',now()),
  ( 'philosophy','Revenge is a confession of pain.',now()),
  ( 'philosophy','Brothers in arms are brothers for life.',now()),
  ( 'philosophy','Fighting a war tests a soldier''s skills, defending his home tests a soldier''s heart.',now()),
  ( 'philosophy','Where there''s a will, there''s a way.',now()),
  ( 'philosophy','A child stolen is a lost hope.',now()),
  ( 'philosophy','The challenge of hope is to overcome corruption.',now()),
  ( 'philosophy','Those who enforce the law must obey the law.',now()),
  ( 'philosophy','The future has many paths -- choose wisely.',now()),
  ( 'philosophy','A failure in planning is a plan for failure.',now()),
  ( 'philosophy','Love comes in all shapes and sizes.',now()),
  ( 'philosophy','Fear is a great motivator.',now()),
  ( 'philosophy','Truth can strike down the specter of fear',now()),
  ( 'philosophy','The swiftest path to destruction is through vengeance.',now()),
  ( 'philosophy','Evil is not born, it is taught.',now()),
  ( 'philosophy','The path to evil may bring great power, but not loyalty.',now()),
  ( 'philosophy','Balance is found in the one who faces his guilt.',now()),
  ( 'philosophy','He who surrenders hope, surrenders life.',now()),
  ( 'philosophy','He who seeks to control fate shall never find peace.',now()),
  ( 'philosophy','Adaptation is the key to survival.',now()),
  ( 'philosophy','Anything that can go wrong will.',now()),
  ( 'philosophy','Without honor, victory is hollow.',now()),
  ( 'philosophy','Without humility, courage is a dangerous game.',now()),
  ( 'philosophy','A great student is what the teacher hopes to be.',now()),
  ( 'philosophy','When destiny calls, the chosen have no choice.',now()),
  ( 'philosophy','Only through fire is a strong sword forged.',now()),
  ( 'philosophy','Crowns are inherited, kingdoms are earned.',now()),
  ( 'philosophy','Who a person truly is cannot be seen with the eye.',now()),
  ( 'philosophy','Understanding is honoring the truth beneath the surface.',now()),
  ( 'philosophy','Who''s the more foolish, the fool or the fool who follows him?',now()),
  ( 'philosophy','The first step toward loyalty is trust.',now()),
  ( 'philosophy','The path of ignorance is guided by fear.',now()),
  ( 'philosophy','The wise man leads, the strong man follows.',now()),
  ( 'philosophy','Our actions define our legacy.',now()),
  ( 'philosophy','Where we are going always reflects where we came from.',now()),
  ( 'philosophy','Those who enslave others inevitably become slaves themselves.',now()),
  ( 'philosophy','Great hope can come from small sacrifices.',now()),
  ( 'philosophy','Friendship shows us who we really are.',now()),
  ( 'philosophy','All warfare is based on deception.',now()),
  ( 'philosophy','Keep your friends close, but keep your enemies closer.',now()),
  ( 'philosophy','The strong survive, the noble overcome.',now()),
  ( 'philosophy','Trust is the greatest of gifts, but it must be earned.',now()),
  ( 'philosophy','One must let go of the past to hold onto the future.',now()),
  ( 'philosophy','Who we are never changes, who we think we are does.',now()),
  ( 'philosophy','A fallen enemy may rise again, but the reconciled one is truly vanquished.',now()),
  ( 'philosophy','The enemy of my enemy is my friend.',now()),
  ( 'philosophy','Strength in character can defeat strength in numbers.',now()),
  ( 'philosophy','Fear is a malleable weapon.',now()),
  ( 'philosophy','To seek something is to believe in its possibility.',now()),
  ( 'philosophy','Struggles often begin and end with the truth.',now()),
  ( 'philosophy','Disobedience is a demand for change.',now()),
  ( 'philosophy','He who faces himself, finds himself.',now()),
  ( 'philosophy','The young are often underestimated.',now()),
  ( 'philosophy','When we rescue others, we rescue ourselves.',now()),
  ( 'philosophy','Choose your enemies wisely, as they may be your last hope.',now()),
  ( 'philosophy','Humility is the only defense against humiliation.',now()),
  ( 'philosophy','When all seems hopeless, a true hero gives hope.',now()),
  ( 'philosophy','A soldier''s most powerful weapon is courage.',now()),
  ( 'philosophy','You must trust in others or success is impossible.',now()),
  ( 'philosophy','One vision can have many interpretations.',now()),
  ( 'philosophy','Alliances can stall true intentions.',now()),
  ( 'philosophy','Morality separates heroes from villains.',now()),
  ( 'philosophy','Sometimes even the smallest doubt can shake the greatest belief.',now()),
  ( 'philosophy','Courage begins by trusting oneself.',now()),
  ( 'philosophy','Never become desperate enough to trust the untrustworthy.',now()),
  ( 'philosophy','Never give up hope, no matter how dark things seem.',now())