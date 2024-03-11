"""create policy table

Revision ID: cfac225b1f8c
Revises:
Create Date: 2024-03-11 15:46:37.435418

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = 'cfac225b1f8c'
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    op.create_table(
        'policies',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('name', sa.String, nullable=False),
        sa.Column('criteria', sa.String, nullable=False),
        sa.Column('value', sa.Integer, nullable=False),
        sa.Column('created_at', sa.DateTime, server_default=sa.text('now()')),
        sa.Column('updated_at', sa.DateTime, server_default=sa.text('now()'), server_onupdate=sa.text('now()')),
    )


def downgrade() -> None:
    op.drop_table('policies')
